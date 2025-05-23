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

	_ = nextLine()
	firstArrayStr := nextLine()
	firstArray := strings.Split(firstArrayStr, separator)
	secondArrayStr := nextLine()
	secondArray := strings.Split(secondArrayStr, separator)

	var array []string
	array = append(array, firstArray...)
	array = append(array, secondArray...)

	for _, v := range array {
		fmt.Printf("%s\n", v)
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// Q
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// 3 行目に M 個の整数 b_1, b_2, ..., b_M が与えられます。
// N 個の整数 a_1, a_2, ..., a_N の後ろに M 個の整数 b_1, b_2, ..., b_M を連結させ、改行区切りで出力してください。

// 入力例 1
// 2 3
// 1 2
// 3 4 5

// 出力例 1
// 1
// 2
// 3
// 4
// 5
