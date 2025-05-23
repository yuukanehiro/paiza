package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a, b, c := nextLine(), nextLine(), nextLine()

	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", c)
}
