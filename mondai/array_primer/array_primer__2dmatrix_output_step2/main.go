package main

import (
	"fmt"
	"strings"
)

func main() {
	a := [][]string{
		{"6", "5", "4", "3", "2", "1"},
		{"3", "1", "8", "8", "1", "3"},
	}

	for i := 0; i < len(a); i++ {
		_s := strings.Join(a[i], " ")
		fmt.Printf("%s\n", _s)
	}
}

// 出力
// 6 5 4 3 2 1
// 3 1 8 8 1 3
