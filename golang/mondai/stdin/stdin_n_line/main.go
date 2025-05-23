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
	var a []string

	for {
		sc.Scan()

		if sc.Text() == "" {
			break
		}

		a = append(a, sc.Text())
	}

	// 最初の行を捨てる
	a = a[1:]

	for _, v := range a {
		fmt.Printf("%s\n", v)
	}
}

// 入力例
// 5
// 4
// 3
// 2
// 1
// 0

// 出力例
// 4
// 3
// 2
// 1
// 0
