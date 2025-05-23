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
	targetNumber := targetInfoArray[1]

	arrayStr := nextLine()
	array := strings.Split(arrayStr, separator)

	fmt.Printf("%s", checkNumber(array, targetNumber))
}

func checkNumber(array []string, targetNumber string) string {
	for _, v := range array {
		if v == targetNumber {
			return "Yes"
		}
	}
	return "No"
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// Q
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数の中に、整数 M が含まれているなら Yes、含まれていないなら No を出力してください。

// 入力例 1
// 5 1
// 1 2 3 4 5

// 出力例 1
// Yes
