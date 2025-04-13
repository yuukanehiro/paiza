package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/no-recover", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			panic("panic in goroutine without recover")
		}()
		w.Write([]byte("Triggered panic without recover\n"))
	})

	http.HandleFunc("/with-recover", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in goroutine:", r)
				}
			}()
			panic("panic in goroutine with recover")
		}()
		w.Write([]byte("Triggered panic with recover\n"))
	})

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
