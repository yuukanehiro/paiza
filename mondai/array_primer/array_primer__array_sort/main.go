package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	const separator = " "
	_ = nextLine() // 一行目はスキップ

	array := nextLineNumbersBySeparator(separator)

	sortAsc(array)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

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

func sortAsc(numbers []int) {
	sort.Ints(numbers)
}

// Q
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数を小さい順にソートし、改行区切りで出力してください。

// 入力例1
// 5
// 5 4 3 2 1

// 出力例1
// 1
// 2
// 3
// 4
// 5
