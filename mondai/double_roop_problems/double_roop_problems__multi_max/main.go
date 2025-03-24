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
	intArrayA := nextLineBySeparator(" ", "int").([]int)
	intArrayB := nextLineBySeparator(" ", "int").([]int)

	var max int = intArrayA[0] * intArrayB[0]
	for i := 0; i < len(intArrayA); i++ {
		for j := 0; j < len(intArrayB); j++ {
			if intArrayA[i]*intArrayB[j] > max {
				max = intArrayA[i] * intArrayB[j]
			}
		}
	}

	fmt.Println(max)
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
// 配列 A と B についての情報が与えられるので、(A の 1 つの要素) × (B の 1 つの要素) の最大値を求めてください。

// Input
// 10 10
// 85 -46 93 44 -40 -75 -75 -18 -54 95
// 1 95 -92 -90 32 -25 36 55 22 86

// Output
// 9025
