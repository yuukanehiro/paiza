package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	_ = nextLine() // 1行目は読み捨て
	max := 5

	for i := 0; i < max; i++ {
		a := nextLine()
		fmt.Printf("%s\n", a)
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 入力例
// 4
// 1 2 3 4
// 0 0 0 0
// 8 1 3 8
// 1 10 100 99
// 15 68 48 15

// 出力例
// 1 2 3 4
// 0 0 0 0
// 8 1 3 8
// 1 10 100 99
// 15 68 48 15
