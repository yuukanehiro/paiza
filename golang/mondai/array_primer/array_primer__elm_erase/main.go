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
	eraseTargetNumber := targetInfoNumbers[1]

	array := nextLineNumbersBySeparator(" ")

	array = append(array[:eraseTargetNumber-1], array[eraseTargetNumber:]...)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func eraseElement(array []int, target int) []int {
	return append(array[:target-1], array[target:]...)
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
// 左から M 番目の要素を削除し、削除した後の N - 1 個の要素を改行区切りで出力してください。
// なお、左端を 1 番目とします。

// 入力例1
// 5 3
// 1 2 3 4 5

// 出力例1
// 1
// 2
// 4
// 5
