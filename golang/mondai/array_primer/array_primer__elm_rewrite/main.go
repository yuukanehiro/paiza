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
	replaceNumber := targetInfoArray[1]

	arrayStr := nextLine()
	array := strings.Split(arrayStr, separator)

	for _, v := range array {
		if v == targetNumber {
			v = replaceNumber
		}
		fmt.Printf("%s\n", v)
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// Q
// 1 行目に整数 A, B, N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、その数が A だった場合、B に書き換えてください。
// 書き換えた N 個の整数を改行区切りで出力してください。

// 入力例 1
// 3 1 5
// 1 2 3 4 5

// 出力例 1
// 1
// 2
// 1
// 4
// 5
