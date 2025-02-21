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
	const targetNum = 1

	_ = nextLine()
	numbers := nextLineBySeparator(" ", "int").([]int)

	for i, v := range numbers {
		if v == targetNum {
			fmt.Printf("%d\n", i+1)
		}
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
// 長さ N の数列 a (a_1, a_2, ..., a_N) が与えられます。
// この数列の何番目に 1 があるか出力してください。
// 数列の 1 つ目の要素を 1 番目とし、数列には必ず 1 がひとつだけ含まれることとします。

// Input
// 5
// 5 4 3 2 1

// Output
// 5
