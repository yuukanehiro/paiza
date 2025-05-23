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
	_ = nextLine() // 1行目は読み飛ばす

	array := nextLineNumbersBySeparator(" ")
	reverse(array)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func reverse(array []int) {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-i-1] = array[len(array)-i-1], array[i]
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

// Q
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数の順番を反転させ、改行区切りで出力してください。

// Input
// 5
// 1 5 2 4 3

// Output
// 3
// 4
// 2
// 5
// 1
