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
	target := infoArray[0]
	divNum := infoArray[1]

	fmt.Println(getDivCount(target, 0, divNum))
}

func getDivCount(n int, count int, divNum int) int {
	if n < divNum {
		return 0
	}

	t := n / divNum
	count = count + 1

	if t < divNum || t%divNum != 0 {
		if count == 0 {
			return 1
		}
		return count
	}

	return getDivCount(t, count, divNum)
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
// 整数 N, M が与えられます。
// N が何回 M で割れるかを求め、出力してください。

// Input
// 16 8

// Output
// 2

// Input
// 81 3

// Output
// 4
