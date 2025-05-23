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

	infoStr := nextLine()
	array := strings.Split(infoStr, separator)
	var firstStr string = array[0]
	var secondStr string = array[1]
	_array := []string{
		secondStr,
		firstStr,
	}

	res := strings.Join(_array, separator)

	fmt.Printf("%s\n", res)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// Q
// 整数 A, B が与えられます。
// A と B の値を入れ替えて、半角スペース区切りで出力してください。

// 入力例 1
// 2 1

// 出力例 1
// 1 2
