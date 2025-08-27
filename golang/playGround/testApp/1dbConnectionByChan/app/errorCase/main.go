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
	req := dbReq{fn: fn, resp: make(chan error, 1)}
	e.ch <- req
	return <-req.resp
}
func (e *DBExecutor) Close() { close(e.ch); <-e.done }

/************** A / B / C **************/
func A(ctx context.Context, exec *DBExecutor) error {
	// 非DBの重い処理は自由に並行OK（ダミー）
	// 3秒待つ
	time.Sleep(3 * time.Second)

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
	// 3秒待つ
	time.Sleep(3 * time.Second)
	return exec.Do(ctx, func(ctx context.Context, conn *sql.Conn) error {
		_, err := conn.ExecContext(ctx, `INSERT INTO demo_seq (step) VALUES ('B');`)
		return err
	})
}

func C(ctx context.Context, exec *DBExecutor) error {
	// 10秒待つ
	time.Sleep(3 * time.Second)
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

// sample-app    | fatal error: all goroutines are asleep - deadlock!
// sample-app    |
// sample-app    | goroutine 1 [sync.WaitGroup.Wait]:
// sample-app    | sync.runtime_SemacquireWaitGroup(0x4000091b38?)
// sample-app    |         /usr/local/go/src/runtime/sema.go:110 +0x2c
// sample-app    | sync.(*WaitGroup).Wait(0x4000098048)
// sample-app    |         /usr/local/go/src/sync/waitgroup.go:118 +0x70
// sample-app    | golang.org/x/sync/errgroup.(*Group).Wait(0x4000098040)
// sample-app    |         /go/pkg/mod/golang.org/x/sync@v0.8.0/errgroup/errgroup.go:56 +0x28
// sample-app    | main.main()
// sample-app    |         /app/main.go:175 +0x5a8
// sample-app    |
// sample-app    | goroutine 7 [select]:
// sample-app    | database/sql.(*DB).connectionOpener(0x40001009c0, {0x2d0750, 0x4000118050})
// sample-app    |         /usr/local/go/src/database/sql/sql.go:1261 +0x80
// sample-app    | created by database/sql.OpenDB in goroutine 1
// sample-app    |         /usr/local/go/src/database/sql/sql.go:841 +0x124
// sample-app    |
// sample-app    | goroutine 12 [select]:
// sample-app    | github.com/go-sql-driver/mysql.(*mysqlConn).startWatcher.func1()
// sample-app    |         /go/pkg/mod/github.com/go-sql-driver/mysql@v1.8.1/connection.go:628 +0x7c
// sample-app    | created by github.com/go-sql-driver/mysql.(*mysqlConn).startWatcher in goroutine 1
// sample-app    |         /go/pkg/mod/github.com/go-sql-driver/mysql@v1.8.1/connection.go:625 +0xf8
// sample-app    |
// sample-app    | goroutine 17 [sync.RWMutex.Lock]:
// sample-app    | sync.runtime_SemacquireRWMutex(0x0?, 0x0?, 0x0?)
// sample-app    |         /usr/local/go/src/runtime/sema.go:105 +0x28
// sample-app    | sync.(*RWMutex).Lock(0x0?)
// sample-app    |         /usr/local/go/src/sync/rwmutex.go:155 +0xfc
// sample-app    | database/sql.(*Conn).close(0x4000098000, {0x0, 0x0})
// sample-app    |         /usr/local/go/src/database/sql/sql.go:2138 +0x98
// sample-app    | database/sql.(*Conn).Close(...)
// sample-app    |         /usr/local/go/src/database/sql/sql.go:2153
// sample-app    | panic({0x22a3e0?, 0x452990?})
// sample-app    |         /usr/local/go/src/runtime/panic.go:792 +0x124
// sample-app    | github.com/go-sql-driver/mysql.(*mysqlConn).watchCancel(0x4000140000, {0x0, 0x0})
// sample-app    |         /go/pkg/mod/github.com/go-sql-driver/mysql@v1.8.1/connection.go:603 +0x2c
// sample-app    | github.com/go-sql-driver/mysql.(*mysqlConn).BeginTx(0x4000140000, {0x0?, 0x0?}, {0xe8dec?, 0x20?})
// sample-app    |         /go/pkg/mod/github.com/go-sql-driver/mysql@v1.8.1/connection.go:490 +0x50
// sample-app    | database/sql.ctxDriverBegin({0x0, 0x0}, 0x0?, {0x2d05b8, 0x4000140000})
// sample-app    |         /usr/local/go/src/database/sql/ctxutil.go:104 +0x84
// sample-app    | database/sql.(*DB).beginDC.func1()
// sample-app    |         /usr/local/go/src/database/sql/sql.go:1906 +0xbc
// sample-app    | database/sql.withLock({0x2cf968, 0x400018c000}, 0x4000040620)
// sample-app    |         /usr/local/go/src/database/sql/sql.go:3572 +0x74
// sample-app    | database/sql.(*DB).beginDC(0x40001009c0, {0x0, 0x0}, 0x400018c000, 0x4000010430, 0x0?)
// sample-app    |         /usr/local/go/src/database/sql/sql.go:1902 +0x84
// sample-app    | database/sql.(*Conn).BeginTx(0x4000098000, {0x0, 0x0}, 0x0)
// sample-app    |         /usr/local/go/src/database/sql/sql.go:2115 +0x6c
// sample-app    | main.A.func1({0x0, 0x0}, 0x0?)
// sample-app    |         /app/main.go:65 +0x3c
// sample-app    | main.NewDBExecutor.func1()
// sample-app    |         /app/main.go:42 +0xbc
// sample-app    | created by main.NewDBExecutor in goroutine 1
// sample-app    |         /app/main.go:38 +0xfc
// sample-app    |
// sample-app    | goroutine 18 [chan receive]:
// sample-app    | main.main.func1()
// sample-app    |         /app/main.go:158 +0x44
// sample-app    | golang.org/x/sync/errgroup.(*Group).Go.func1()
// sample-app    |         /go/pkg/mod/golang.org/x/sync@v0.8.0/errgroup/errgroup.go:78 +0x54
// sample-app    | created by golang.org/x/sync/errgroup.(*Group).Go in goroutine 1
// sample-app    |         /go/pkg/mod/golang.org/x/sync@v0.8.0/errgroup/errgroup.go:75 +0x94
// sample-app    |
// sample-app    | goroutine 19 [chan receive]:
// sample-app    | main.(*DBExecutor).Do(...)
// sample-app    |         /app/main.go:53
// sample-app    | main.A({0x0?, 0x0?}, 0x400009c000)
// sample-app    |         /app/main.go:64 +0x70
// sample-app    | main.main.func2()
// sample-app    |         /app/main.go:165 +0x30
// sample-app    | golang.org/x/sync/errgroup.(*Group).Go.func1()
// sample-app    |         /go/pkg/mod/golang.org/x/sync@v0.8.0/errgroup/errgroup.go:78 +0x54
// sample-app    | created by golang.org/x/sync/errgroup.(*Group).Go in goroutine 1
// sample-app    |         /go/pkg/mod/golang.org/x/sync@v0.8.0/errgroup/errgroup.go:75 +0x94
// sample-app    |
// sample-app    | goroutine 20 [chan receive]:
// sample-app    | main.main.func3()
// sample-app    |         /app/main.go:171 +0x3c
// sample-app    | golang.org/x/sync/errgroup.(*Group).Go.func1()
// sample-app    |         /go/pkg/mod/golang.org/x/sync@v0.8.0/errgroup/errgroup.go:78 +0x54
// sample-app    | created by golang.org/x/sync/errgroup.(*Group).Go in goroutine 1
// sample-app    |         /go/pkg/mod/golang.org/x/sync@v0.8.0/errgroup/errgroup.go:75 +0x94
// sample-app    | exit status 2
// sample-app exited with code 1
