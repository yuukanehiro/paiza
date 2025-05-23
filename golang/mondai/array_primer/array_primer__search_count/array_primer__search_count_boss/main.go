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

	arrayStr := nextLine()
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
// 1 行目に整数 N, M が与えられます。
// 2 行目に M 個の整数 a_1, a_2, ..., a_M が与えられます。
// 以下の形式で標準入力によって与えられます。

// N M
// a_1 a_2 ... a_M

// M 個の整数に N が何個あるか数え、出力してください。
// また、末尾に改行を入れ、余計な文字、空行を含んではいけません。

// input
// 1 5
// 1 1 1 2 2

// output
// 3
