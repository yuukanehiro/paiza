package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
)

/************** DB専用ワーカー（アクター） **************/
type dbReq struct {
	ctx  context.Context
	fn   func(ctx context.Context, conn *sql.Conn) error
	resp chan error
}

type DBExecutor struct {
	ch   chan dbReq
	done chan struct{}
}

func NewDBExecutor(ctx context.Context, db *sql.DB) (*DBExecutor, error) {
	// conn を取るときの ctx は “確立まで” にだけ使う。以降は各リクエストの ctx を使う
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	ex := &DBExecutor{
		ch:   make(chan dbReq),
		done: make(chan struct{}),
	}
	go func() {
		defer close(ex.done)
		defer conn.Close()
		for req := range ex.ch {
			err := req.fn(req.ctx, conn)
			req.resp <- err
			close(req.resp)
		}
	}()
	return ex, nil
}

func (e *DBExecutor) Do(ctx context.Context, fn func(ctx context.Context, conn *sql.Conn) error) error {
	req := dbReq{ctx: ctx, fn: fn, resp: make(chan error, 1)}
	e.ch <- req
	return <-req.resp
}
func (e *DBExecutor) Close() { close(e.ch); <-e.done }

/************** A / B / C **************/
func A(ctx context.Context, exec *DBExecutor) error {
	// 非DBの重い処理は自由に並行OK（ダミー）
	// 1秒待つ
	time.Sleep(1 * time.Second)

	// DB区間
	return exec.Do(ctx, func(ctx context.Context, conn *sql.Conn) error {
		tx, err := conn.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `
			CREATE TABLE IF NOT EXISTS demo_seq (
			  id INT AUTO_INCREMENT PRIMARY KEY,
			  step VARCHAR(16) NOT NULL,
			  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		); err != nil {
			_ = tx.Rollback()
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO demo_seq (step) VALUES ('A');`); err != nil {
			_ = tx.Rollback()
			return err
		}
		return tx.Commit()
	})
}

func B(ctx context.Context, exec *DBExecutor) error {
	// 1秒待つ
	time.Sleep(1 * time.Second)
	return exec.Do(ctx, func(ctx context.Context, conn *sql.Conn) error {
		_, err := conn.ExecContext(ctx, `INSERT INTO demo_seq (step) VALUES ('B');`)
		return err
	})
}

func C(ctx context.Context, exec *DBExecutor) error {
	// 1秒待つ
	time.Sleep(1 * time.Second)
	return exec.Do(ctx, func(ctx context.Context, conn *sql.Conn) error {
		_, err := conn.ExecContext(ctx, `INSERT INTO demo_seq (step) VALUES ('C');`)
		return err
	})
}

/************** エントリポイント **************/
func main() {
	// ---- 環境変数から DSN 構築 ----
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "3306")
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASS", "password")
	name := getenv("DB_NAME", "testdb")

	timeoutSec, _ := strconv.Atoi(getenv("DB_CONN_TIMEOUT_SEC", "15"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	cfg := mysql.Config{
		User:                 user,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", host, port),
		DBName:               name,
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.FixedZone("Asia/Tokyo", 9*60*60),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認（リトライ軽め）
	if err := pingWithRetry(ctx, db, 10, 500*time.Millisecond); err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}

	// 単一コネクションのDBワーカー
	exec, err := NewDBExecutor(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	defer exec.Close()

	// A→B→C の順序制約（シンプルな通知チャネル）
	aDone := make(chan struct{})
	bDone := make(chan struct{})

	var g errgroup.Group

	// A→B→C の順序で実行
	// 排他制御が行われているか確認するために敢えてB, A, Cの順で起動させている
	// 結果としては排他制御が行われて A, B, C の順で実行される
	// B
	g.Go(func() error {
		<-aDone
		err := B(ctx, exec)
		close(bDone)
		return err
	})
	// A
	g.Go(func() error {
		err := A(ctx, exec)
		close(aDone)
		return err
	})
	// C
	g.Go(func() error {
		<-bDone
		return C(ctx, exec)
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("A/B/C failed: %v", err)
	}

	// 結果表示
	rows, err := db.QueryContext(ctx, `SELECT id, step, created_at FROM demo_seq ORDER BY id;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("== demo_seq ==")
	for rows.Next() {
		var id int
		var step string
		var ts time.Time
		if err := rows.Scan(&id, &step, &ts); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s at %s\n", id, step, ts.Format(time.RFC3339))
	}
}

func pingWithRetry(ctx context.Context, db *sql.DB, retries int, interval time.Duration) error {
	for i := 0; i < retries; i++ {
		if err := db.PingContext(ctx); err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}
	}
	return fmt.Errorf("ping failed after %d retries", retries)
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
