package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("start")

	pr, pw := io.Pipe()

	// Writer goroutine
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered in writer:", r)
			}
		}()

		writer := csv.NewWriter(pw)
		_ = writer.Write([]string{"hello", "world"})
		writer.Flush()

		fmt.Println("writer: done writing CSV")

		// NOTE: 故意に Close を呼ばない！
		// pw.Close()
		// panic("intentional panic before Close()") // これでも同じ
	}()

	// Reader (main goroutine)
	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("start io.Copy... (this will hang)")

	_, err = io.Copy(file, pr)
	if err != nil {
		fmt.Println("io.Copy error:", err)
	} else {
		fmt.Println("io.Copy success")
	}

	fmt.Println("done")
}
