package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	a := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}

	// 各行の要素数を出力
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\n", len(a[i]))
	}
}

// 出力
// 1
// 2
// 3
