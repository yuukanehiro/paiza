package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	const separator = " "

	targetInfoStr := nextLine()
	targetInfoArray := strings.Split(targetInfoStr, separator)
	numberFromIndex, _ := strconv.Atoi(targetInfoArray[0])
	numberToIndex, _ := strconv.Atoi(targetInfoArray[1])

	arrayStr := nextLine()
	array := strings.Split(arrayStr, separator)

	_array := array[numberFromIndex-1 : numberToIndex]

	for _, v := range _array {
		fmt.Printf("%s\n", v)
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// Q
// 1 行目に整数 A, B, N （A ≦ B）が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、左から A 番目から B 番目までの数を抜き出し、改行区切りで出力してください。
// なお、左端を 1 番目とします。

// 入力例 1
// 2 4 5
// 1 2 3 4 5

// 出力例 1
// 2
// 3
// 4
