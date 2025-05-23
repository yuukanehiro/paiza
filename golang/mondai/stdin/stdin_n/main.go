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
	_ = nextLine()
	a := nextLine()

	var b []string
	b = strings.Split(a, " ")

	for _, v := range b {
		fmt.Printf("%s\n", v)
	}
}
