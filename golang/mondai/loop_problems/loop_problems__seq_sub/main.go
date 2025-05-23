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
	infoArray := nextLineBySeparator(" ", "int").([]int)
	max := infoArray[0]

	arrayA := nextLineBySeparator(" ", "int").([]int)
	arrayB := nextLineBySeparator(" ", "int").([]int)

	if len(arrayA) != len(arrayB) {
		fmt.Println("Input error")
		return
	}

	for i := 0; i < max; i++ {
		fmt.Println(arrayA[i] - arrayB[i])
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
// 長さ N の数列 a (a_1, a_2, ..., a_N) と b (b_1, b_2, ..., b_N) が与えられます。
// a の各要素から b の各要素を引いた結果 (a_1 - b_1, a_2 - b_2, ..., a_N - b_N) を、改行区切りで出力してください。

// Input
// 5
// 1 2 3 4 5
// 5 4 3 2 1

// Output
// -4
// -2
// 0
// 2
// 4
