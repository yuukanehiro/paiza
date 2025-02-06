package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
	}

	// 要素数を出力
	var c int
	for i := 0; i < len(a); i++ {
		c += len(a[i])
	}

	fmt.Printf("%d\n", c)
}

// 入力例
// 1 2 3 4 5 6
// 8 1 3 3 1 8

// 出力例
// 12
