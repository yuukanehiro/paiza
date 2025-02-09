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

	var count int = 0
	for _, v := range array {
		count++
		if v == targetNumber {
			break
		}
	}

	fmt.Printf("%d\n", count)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// Q
// 1 行目に整数 N, M が与えられます。
// 2 行目に M 個の整数 a_1, a_2, ..., a_M が与えられます。
// 整数 N が、M 個の整数の左から何番目にあるか出力してください。
// 左端を 1 番目とし、N は M 個の整数に必ず 1 つだけ含まれるものとします。

// input
// 3 5
// 1 2 3 4 5

// output
// 3
