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

	targetInfoStr := nextLine()
	targetInfoArray := strings.Split(targetInfoStr, separator)
	targetNumber := targetInfoArray[0]

	arrayStr := "1 2 5 1 4 3 2 5 1 4"
	array := strings.Split(arrayStr, separator)

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
// 整数 N が与えられます。
// 以下の配列に含まれる N の個数を出力してください。
// また、N は以下の配列に 1 個以上含まれるものとします。

// 1 2 5 1 4 3 2 5 1 4

// 入力例1
// 1

// 出力例1
// 3
