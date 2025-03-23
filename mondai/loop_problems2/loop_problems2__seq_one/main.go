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
	var targetNum = 1
	var targetCountArr []int
	for i, v := range numbers {
		// n番目
		n := i + 1
		if v == targetNum {
			targetCountArr = append(targetCountArr, n)
		}
	}

	for _, v := range targetCountArr {
		fmt.Println(v)
	}
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
// a_1, a_2, ..., a_N のうち、1 がある位置を先頭から順に改行区切りで出力してください。
// a_1 を 1 番目とし、a_1, a_2, ..., a_N には少なくとも 1 個は 1 が含まれます。

// Input
// 5
// 5 3 1 3 5

// Output
// 3
