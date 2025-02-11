package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	targetStr := "1 5 9 7 3 2 4 8 6 10"
	separator := " "
	array := strings.Split(targetStr, separator)
	targetNumber := nextLine()

	fmt.Printf("%d\n", countNumber(array, targetNumber))
}

func countNumber(array []string, targetNumber string) int {
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
// 整数 N が与えられます。
// 整数 N が、以下の配列の左から何番目にあるか出力してください。
// 左端を 1 番目とし、N は以下の配列に必ず含まれるものとします。

// 1 5 9 7 3 2 4 8 6 10

// 入力例1
// 5

// 出力例1
// 2
