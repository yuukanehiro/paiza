package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	const separator = " "
	const targetNumber = "1"

	targetStr := "1 2 2 1 2 1 2 1 1 1"

	array := strings.Split(targetStr, separator)

	fmt.Printf("%d\n", countNumber(array, targetNumber))
}

func countNumber(array []string, targetNumber string) int {
	var count int = 0
	for _, v := range array {
		if v == targetNumber {
			count++
		}
	}

	return count
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// Q
// 以下のような配列があります。
// 1 2 2 1 2 1 2 1 1 1
// この中に含まれる 1 の個数を出力してください。

// Output
// 6
