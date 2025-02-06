package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{1, 2, 3},
		{8, 1, 3},
		{10, 100, 1},
	}

	//  2 行目 3 列目の要素を出力
	fmt.Printf("%d\n", a[1][2]) // 3
}
