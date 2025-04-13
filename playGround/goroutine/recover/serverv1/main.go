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
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered in writer:", r)

					// pr にエラー通知（これで io.Copy を unblock）
					func() {
						defer func() {
							if r := recover(); r != nil {
								fmt.Println("recovered inside pw.CloseWithError():", r)
							}
						}()
						err := pw.CloseWithError(fmt.Errorf("writer failed: %v", r))
						if err != nil {
							fmt.Println("pw.CloseWithError() error:", err)
						} else {
							fmt.Println("pw.CloseWithError() success")
						}
					}()
				}
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

			// 強制 panic → recover → CloseWithError
			panic("panic after writing CSV to pipe")
		}()

		file, err := os.Create("output.csv")
		if err != nil {
			http.Error(w, fmt.Sprintf("file create error: %v", err), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, pr)
		if err != nil {
			// io.Copy は CloseWithError のエラーを受け取る
			fmt.Println("io.Copy error:", err)
		} else {
			fmt.Println("io.Copy success")
		}

		fmt.Println("Pipe data written to output.csv")

		w.Write([]byte("done\n"))
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
