package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a := nextLine()
	_a, _ := strconv.Atoi(a)

	s := nextLine()
	var b []string
	b = strings.Split(s, " ")

	fmt.Printf("%s\n", b[_a-1])
}

// 入力例
// 5
// 8 1 3 1 3 8 3 8 1 1

// 出力例
// 3
