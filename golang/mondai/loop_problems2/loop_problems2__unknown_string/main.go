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
	const endMark = "EOF"

	infoArray := nextLineBySeparator(" ", "string").([]string)

	var array []string
	for _, v := range infoArray {
		array = append(array, v)
		if v == endMark {
			break
		}
	}

	for _, v := range array {
		fmt.Printf("%s\n", v)
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
// 複数の文字列が入力されます。文字列の数はわかりません。
// EOF が入力されるまで、受け取った文字列を改行区切りで出力してください。

// Input
// abc def ghi jkl EOF

// Output
// abc
// def
// ghi
// jkl
// EOF
