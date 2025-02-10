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
	targetInfoNumbers := nextLineNumbersBySeparator(" ")
	pushNumber := targetInfoNumbers[1]

	array := nextLineNumbersBySeparator(" ")
	array = append(array, pushNumber)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
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
// N 個の整数の末尾に整数 M を挿入し、改行区切りで出力してください。

// 入力例1
// 5 1
// 1 2 3 4 5

// 出力例1
// 1
// 2
// 3
// 4
// 5
// 1
