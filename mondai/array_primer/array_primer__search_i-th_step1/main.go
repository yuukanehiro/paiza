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
	const targetNumber = "8"

	arrayStr := "1 10 2 9 3 8 4 7 5 6"
	array := strings.Split(arrayStr, separator)

	fmt.Printf("%d\n", checkNumber(array, targetNumber))
}

func checkNumber(array []string, targetNumber string) int {
	var count int = 0
	for _, v := range array {
		count += 1
		if v == targetNumber {
			break
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

// 1 10 2 9 3 8 4 7 5 6

// この配列の中で、8 が左から何番目にあるか出力してください。
// 左端を 1 番目とします。

// Output
// 6
