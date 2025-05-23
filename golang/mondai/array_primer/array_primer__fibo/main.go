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
	targetCount := infoArray[0]

	array := getFibonacciArray(targetCount)
	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func getFibonacciArray(n int) []int {
	var fibonacciArray []int

	if n == 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	fibonacciArray = []int{0, 1}

	for i := 2; i < n; i++ {
		tmp := fibonacciArray[i-2] + fibonacciArray[i-1]
		fibonacciArray = append(fibonacciArray, tmp)
	}

	return fibonacciArray
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
// N 番目までのフィボナッチ数を出力してください。

// フィボナッチ数は

// F_0 = 0
// F_1 = 1
// F_(n+2) = F_n + F_(n+1) (n は 0 以上)

// とし、F_0 を 1 番目とします。

// Input
// 3

// Output
// 0
// 1
// 1
