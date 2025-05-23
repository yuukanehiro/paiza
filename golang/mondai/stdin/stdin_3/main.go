package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a := nextLine()
	var b []string
	b = strings.Split(a, " ")

	for _, v := range b {
		fmt.Printf("%s\n", v)
	}
}

// 入力例
// a a a

// 出力例
// a
// a
// a
