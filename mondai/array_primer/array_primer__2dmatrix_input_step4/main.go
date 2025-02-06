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
	a := nextLine()
	arr := strings.Split(a, " ")
	max := arr[0] // 行数取得
	_max, _ := strconv.Atoi(max)

	for i := 0; i < _max; i++ {
		b := nextLine()
		fmt.Printf("%s\n", b)
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 入力例
// 4 3
// 1 2 3
// 8 1 3
// 10 100 0
// 12 24 84

// 出力例
// 1 2 3
// 8 1 3
// 10 100 0
// 12 24 84
