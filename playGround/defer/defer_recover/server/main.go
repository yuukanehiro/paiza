package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start /generate")

		pr, pw := io.Pipe()

		go func() {
			// Close B（recover）
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered in writer:", r)

					err := pw.Close()
					// ⭐️Close AでPanicする状況の場合、Close Bでもpanicしたらまたpanicする可能性高いかも？
					// ⭐️その場合はrecover()が再度行われずサーバがダウンします。
					// panic("panic after pw.Close() B")
					if err != nil {
						fmt.Println("B: pw.Close() error:", err)
					} else {
						fmt.Println("B: pw.Close() success")
					}
					fmt.Println("after pw.Close() B")
				}
			}()

			// Close A（通常パス）
			defer func() {
				// ⭐️Close Aでpanicする場合
				err := pw.Close()
				// panic("panic after pw.Close() A")
				if err != nil {
					fmt.Println("A: pw.Close() error:", err)
				} else {
					fmt.Println("A: pw.Close() success")
				}
				fmt.Println("after pw.Close() A")
			}()

			defer fmt.Println("defer 1")
			defer fmt.Println("defer 2")
			defer fmt.Println("defer 3")

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

			// 強制 panic
			panic("panic after writing CSV to pipe")
		}()

		// ファイル作成
		file, err := os.Create("output.csv")
		if err != nil {
			http.Error(w, fmt.Sprintf("file create error: %v", err), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// ファイルへ書き込み
		_, err = io.Copy(file, pr)
		if err != nil {
			fmt.Println("io.Copy error:", err)
		}

		fmt.Println("Pipe data written to output.csv")
		w.Write([]byte("done\n"))
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// ====
// Server listening on :8080
// start /generate
// progress: 0%
// progress: 10%
// progress: 20%
// progress: 30%
// progress: 40%
// progress: 50%
// progress: 60%
// progress: 70%
// progress: 80%
// progress: 90%
// progress: 100%
// CSV data written to pipe
// defer 3
// defer 2
// defer 1
// A: pw.Close() success
// after pw.Close() A
// recovered in writer: panic after writing CSV to pipe
// B: pw.Close() success // ⭐️pw.Close() Aがあるからpw.Close() Bは不要かもという確認でした。意図したものだということで承知致しました。ご確認有難う御座います。
// after pw.Close() B
// Pipe data written to output.csv
