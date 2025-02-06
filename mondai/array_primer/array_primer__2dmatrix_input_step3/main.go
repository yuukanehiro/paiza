package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	lineCount := nextLine() // 行数取得
	max, _ := strconv.Atoi(lineCount)

	for i := 0; i < max; i++ {
		line := nextLine() // 行取得
		fmt.Printf("%s\n", line)
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 入力
// 3
// 8 1 3 8 1
// 1 5 6 4 7
// 1 100 56 25 15

// 出力
// 8 1 3 8 1
// 1 5 6 4 7
// 1 100 56 25 15
