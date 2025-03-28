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

	var lineSums []int
	for i := 0; i < len(numbers); i++ {
		sum := 0
		for j := 0; j < len(numbers[i]); j++ {
			sum += numbers[i][j]
		}
		lineSums = append(lineSums, sum)
	}

	for _, v := range lineSums {
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
// 整数 N , K と N 行 K 列 の二次元配列 A が与えられるので、 A の行ごとの和を出力してください。

// Input
// 3 3
// 1 2 3
// 4 5 6
// 7 8 9

// Output
// 6
// 15
// 24
