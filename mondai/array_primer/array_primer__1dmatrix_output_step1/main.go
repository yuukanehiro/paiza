package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	a := "5 1 3 4 5 12 6 8 1 3"

	var b []string
	b = strings.Split(a, " ")

	fmt.Printf("%d\n", len(b)) // 10
}
