package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	const separator = " "

	line := "5 12 6 84 14 25 44 3 7 20"
	targetNumber := nextLine()
	array := strings.Split(line, separator)

	fmt.Printf("%s\n", checkArray(array, targetNumber))
}

func checkArray(array []string, targetNumber string) string {
	for _, v := range array {
		if v == targetNumber {
			return "Yes"
		}
	}
	return "No"
}

// Q
// 整数 N が与えられます。
// 以下の形式で標準入力によって与えられます。
// N
// 入力値最終行の末尾に改行が１つ入ります。

// 入力例1
// 12

// 出力例1
// Yes
