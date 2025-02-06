package main

import (
	"bufio"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
