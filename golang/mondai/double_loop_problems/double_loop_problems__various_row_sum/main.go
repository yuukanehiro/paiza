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
	lineCount := infoArray[0]

	var numbers [][]int
	for i := 0; i < lineCount; i++ {
		numbers = append(numbers, nextLineBySeparator(" ", "int").([]int))
	}

	var sums []int
	for i := 0; i < len(numbers); i++ {
		var sum int = 0
		for j := 0; j < len(numbers[i]); j++ {
			if j == 0 {
				continue
			}
			sum += numbers[i][j]
		}
		sums = append(sums, sum)
	}

	for _, v := range sums {
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
// 整数 N が与えられるので、次の処理を N 回してください。
// ・ 配列のサイズ K とその要素 A1 ... AK が与えられるので、全ての要素の和を求めて出力してください。

// Input
// 3
// 3 1 2 3
// 3 4 5 6
// 3 7 8 9

// Output
// 6
// 15
// 24
