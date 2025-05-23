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
	_ = nextLine()
	numbers := nextLineBySeparator(" ", "int").([]int)

	// numbersの中から各要素とIndex+1の和を計算し、最大値を求める
	var maxSum int = 0
	for i, v := range numbers {
		// n番目
		n := i + 1
		sum := n + v
		if sum > maxSum {
			maxSum = sum
		}
	}

	fmt.Println(maxSum)
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 行を取得してinterface{}で返却
// 注意: interface{}で返却されるので、キャストを忘れないこと
// 利用例
// ・array := nextLineBySeparator(" ", "string").([]string) // []stringで取得
// ・array := nextLineBySeparator(" ", "int").([]int)　// []intで取得
func nextLineBySeparator(separator string, elementType string) interface{} {
	line := nextLine()

	var numberArray []string
	numberArray = strings.Split(line, separator)

	if len(numberArray) == 0 {
		return nil
	}

	if elementType == "string" {
		return numberArray
	} else if elementType == "int" {

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
	} else {
		fmt.Println("elementType error")
		return nil
	}
}

// Q
// N 個の整数 a_1, a_2, ..., a_N が与えられます。
// a_i に i を足したとき、a_1 , ... , a_N の最大値を出力してください。

// input
// 5
// 1 2 3 4 5

// output
// 10
