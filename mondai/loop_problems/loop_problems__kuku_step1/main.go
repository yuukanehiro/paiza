package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var targetKuKUNumberEight int = 8

	var resArrayStr []string
	for i := 1; i <= 9; i++ {
		n := targetKuKUNumberEight * i
		s := strconv.Itoa(n)

		resArrayStr = append(resArrayStr, s)
	}

	resStr := strings.Join(resArrayStr, " ")
	fmt.Println(resStr)
}

// Q
// 九九の 8 の段を半角スペース区切りで出力してください。

// Output
// 8 16 24 32 40 48 56 64 72
