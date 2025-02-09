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
	max := b[0]
	_max, _ := strconv.Atoi(max)
	rowCount := b[2]
	_rowCount, _ := strconv.Atoi(rowCount)
	colCount := b[3]
	_colCount, _ := strconv.Atoi(colCount)

	var arr [][]string
	for i := 0; i < _max; i++ {
		tmp := nextLine()
		arr = append(arr, strings.Split(tmp, separator))
	}

	fmt.Printf("%s", arr[_rowCount-1][_colCount-1])
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 問題
// 1 行目に整数 N, M, K, L が与えられます。
// 2 行目以降に N 行 M 列の二次元配列が与えられます。
// 配列の K 行目 L 列目の要素を出力してください。
// 上から i 番目、左から j 番目の整数は a_ij です。

// input
// 4 3 2 1
// 1 2 3
// 10 100 0
// 8 1 3
// 7 21 54

// output
// 10
