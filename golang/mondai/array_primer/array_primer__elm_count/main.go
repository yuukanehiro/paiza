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
	targetInfos := nextLineNumbersBySeparator(" ")
	targetNumber := targetInfos[1]
	array := nextLineNumbersBySeparator(" ")

	var count int
	for _, v := range array {
		if v == targetNumber {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 行を取得して[]intに変換して返却
func nextLineNumbersBySeparator(separator string) []int {
	line := nextLine()

	var numberArray []string
	numberArray = strings.Split(line, separator)

	if len(numberArray) == 0 {
		return nil
	}

	var numbers []int
	for _, v := range numberArray {
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("input error", err)
			return nil
		}

		numbers = append(numbers, number)
	}

	return numbers
}

// Q
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数に含まれている M の個数を出力してください。

// 入力例 1
// 5 1
// 1 1 1 2 2

// 出力例 1
// 3
