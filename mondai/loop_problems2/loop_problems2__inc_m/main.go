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
	addNumber := infoArray[1]
	limit := infoArray[2]

	var count int = 0
	if target >= limit {
		fmt.Printf("%d\n", count)
		return
	}

	for i := 0; target <= limit; i++ {
		target += addNumber
		count += 1
	}

	fmt.Printf("%d\n", count)
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
// 整数 N, M, K が与えられます。
// N が M ずつ増えるとき、何回目に K を越えるか出力してください。

// Input1
// 1 1 10

// Output1
// 10

// Input2
// 10 1 1

// Output2
// 0
