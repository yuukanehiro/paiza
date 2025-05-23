package main

import "fmt"

func main() {
	const requestCount = 8
	const requestNumber = 3

	for i := 0; i < requestCount; i++ {
		fmt.Printf("%d\n", requestNumber)
	}
}

// Q
// 3 を 8 回、改行区切りで出力してください。

// Output
// 3
// 3
// 3
// 3
// 3
// 3
// 3
// 3
