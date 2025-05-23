package main

import (
	"fmt"
	"strings"
)

func main() {
	a := [][]string{
		{"1", "3", "5", "7"},
		{"8", "1", "3", "8"},
	}

	for _, v := range a {
		_v := strings.Join(v, " ")
		fmt.Printf("%s\n", _v)
	}
}

// 出力
// 1 3 5 7
// 8 1 3 8
