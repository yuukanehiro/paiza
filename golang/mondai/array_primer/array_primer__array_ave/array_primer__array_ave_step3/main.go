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
	_ = nextLine() // 1行目は不要なので読み込みだけして捨てる

	array := nextLineBySeparator(" ", "int").([]int)

	average := getAverage(array)

	var result []int
	for _, v := range array {
		if float64(v) >= average {
			result = append(result, v)
		}
	}

	for _, v := range result {
		fmt.Printf("%d\n", v)
	}
}

func getAverage(array []int) float64 {
	if len(array) == 0 {
		return 0.0
	}

	sum := 0
	for _, v := range array {
		sum += v
	}

	return float64(sum) / float64(len(array))
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
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、N 個の整数の平均以上の数をすべて、入力された順に改行区切りで出力してください。

// Input
// 5
// 1 2 3 4 5

// Output
// 3
// 4
// 5