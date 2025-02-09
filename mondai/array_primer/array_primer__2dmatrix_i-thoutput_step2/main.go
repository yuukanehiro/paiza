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

	a := nextLine()
	b := strings.Split(a, separator)
	rowCount := b[0]
	_rowCount, _ := strconv.Atoi(rowCount)
	colCount := b[1]
	_colCount, _ := strconv.Atoi(colCount)

	str := [][]string{
		func() []string {
			tmp := nextLine()
			return strings.Split(tmp, separator)
		}(),
		func() []string {
			tmp := nextLine()
			return strings.Split(tmp, separator)
		}(),
		func() []string {
			tmp := nextLine()
			return strings.Split(tmp, separator)
		}(),
	}

	fmt.Printf("%s", str[_rowCount-1][_colCount-1])
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 入力
// 2 3 // 行数 列数
// 1 2 3
// 4 5 6
// 7 8 9

// 出力
// 6
