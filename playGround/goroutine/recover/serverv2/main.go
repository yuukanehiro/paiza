package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start /generate")

		pr, pw := io.Pipe()
		errCh := make(chan error, 1) // 書き手のエラーを伝えるチャネル

		// バックアップクローザー（保険）
		go func() {
			<-time.After(5 * time.Second)
			errCh <- fmt.Errorf("timeout: writer did not close")
			_ = pw.CloseWithError(fmt.Errorf("forced close after timeout"))
		}()

		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered in writer:", r)
					err := pw.CloseWithError(fmt.Errorf("writer panic: %v", r))
					errCh <- err
				}
			}()

			defer func() {
				if err := pw.Close(); err != nil {
					fmt.Println("pw.Close() error:", err)
					errCh <- err
				} else {
					errCh <- nil
				}
			}()

			writer := csv.NewWriter(pw)
			for i := 0; i <= 100; i++ {
				record := []string{fmt.Sprintf("Row %d", i), fmt.Sprintf("%d", i*i)}
				if err := writer.Write(record); err != nil {
					panic(err)
				}
				if i%10 == 0 {
					fmt.Printf("progress: %d%%\n", i)
				}
			}
			writer.Flush()

			fmt.Println("CSV data written to pipe")

			// ※ panic をコメントアウトすれば正常パスもテスト可
			panic("test panic")
		}()

		// メイン側：CSVをファイルに書き出し
		file, err := os.Create("output.csv")
		if err != nil {
			http.Error(w, fmt.Sprintf("file create error: %v", err), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, pr)
		if err != nil {
			fmt.Println("io.Copy error:", err)
		}

		// クローズエラーがあれば HTTP 500
		if closeErr := <-errCh; closeErr != nil {
			fmt.Println("writer close error:", closeErr)
			http.Error(w, "internal error: writer failed to close", http.StatusInternalServerError)
			return
		}

		fmt.Println("Pipe data written to output.csv")
		w.Write([]byte("done\n"))
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
