package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	a := "5 1 3 4 5 12 6 8 1 3"

	var b []string
	b = strings.Split(a, " ")

	c := b[3]
	// stringをintに変換
	_c, _ := strconv.Atoi(c)

	fmt.Printf("%d\n", int(_c)) // 4
}