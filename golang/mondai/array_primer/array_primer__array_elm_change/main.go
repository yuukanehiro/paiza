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
	numberA, _ := strconv.Atoi(targetInfoArray[0])
	numberB, _ := strconv.Atoi(targetInfoArray[1])

	arrayStr := nextLine()
	array := strings.Split(arrayStr, separator)

	array[numberB-1], array[numberA-1] = array[numberA-1], array[numberB-1]

	for _, v := range array {
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
// N 個の整数の左から A 番目の数と B 番目の数の値を入れ替えて、改行区切りで出力してください。
// なお、左端を 1 番目とします。

// 入力例 1
// 2 3 5
// 1 2 3 4 5

// 出力例 1
// 1
// 3
// 2
// 4
// 5
