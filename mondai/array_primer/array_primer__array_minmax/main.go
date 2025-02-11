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
	max := getMax(array)
	min := getMin(array)

	fmt.Printf("%d %d\n", max, min)
}

// 配列の最大値を取得
func getMax(array []int) int {
	max := array[0]

	for _, v := range array {
		if v > max {
			max = v
		}
	}

	return max
}

// 配列の最小値を取得
func getMin(array []int) int {
	min := array[0]

	for _, v := range array {
		if v < min{
			min = v
		}
	}

	return min
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
// N 個の整数のうち、最大の数と最小の数を半角スペース区切りで出力してください。
// N 個の整数を大きい順や小さい順に並び替える操作を考えて解いてみましょう。

// Input
// 5
// 1 3 5 2 4

// Output
// 5 1