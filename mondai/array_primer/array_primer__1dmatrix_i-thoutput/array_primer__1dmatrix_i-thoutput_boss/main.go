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
	a_array := strings.Split(a, " ")
	i := a_array[1] // a_array[1]を取り出す
	_i, _ := strconv.Atoi(i)

	s := nextLine()
	var b []string
	b = strings.Split(s, " ")

	fmt.Printf("%s\n", b[_i-1])
}

// 入力例
// 8 3
// 3 1 8 1 3 8 8 1

// 出力例
// 8
