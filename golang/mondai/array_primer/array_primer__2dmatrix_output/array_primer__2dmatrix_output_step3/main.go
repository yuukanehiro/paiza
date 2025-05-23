package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	a := [][]int{
		{1, 2, 3, 4},
		{6, 5, 4, 3},
		{3, 1, 8, 1},
	}

	// 列数を出力
	fmt.Printf("%d\n", len(a[0])) // 4
}
