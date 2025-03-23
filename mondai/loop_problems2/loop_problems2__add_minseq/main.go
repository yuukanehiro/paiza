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

	min := numbers[0] + 1
	for i, v := range numbers {
		// n番目
		n := i + 1
		sum := v + n
		if sum < min {
			min = sum
		}
	}

	fmt.Println(min)
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
// a_i に i を足したとき、N 個の整数の最小値を出力してください。

// Input1
// 5
// 1 2 3 4 5

// Output1
// 2
