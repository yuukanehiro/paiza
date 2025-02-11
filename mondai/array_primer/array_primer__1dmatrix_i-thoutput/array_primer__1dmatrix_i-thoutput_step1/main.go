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

	s := "1 3 5 4 6 2 1 7 1 5"
	var b []string
	b = strings.Split(s, " ")

	fmt.Printf("%s\n", b[_a-1])
}

// 入力例
// 3

// 出力例
// 5
