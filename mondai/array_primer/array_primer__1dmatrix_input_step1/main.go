package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	a := "8 1 3 3 8 1 1 3 8 8"

	var b []string
	b = strings.Split(a, " ")

	for _, v := range b {
		fmt.Printf("%s\n", v)
	}
}

// 出力
// 8
// 1
// 3
// 3
// 8
// 1
// 1
// 3
// 8
// 8
