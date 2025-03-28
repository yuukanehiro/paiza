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
	const separator = " "
	infoArray := nextLineBySeparator(separator, "int").([]int)
	lineCount := infoArray[0]

	for i := 0; i < lineCount; i++ {
		var numbers []string
		for j := 0; j < i+1; j++ {
			numbers = append(numbers, strconv.Itoa(j+1))
		}
		fmt.Println(strings.Join(numbers, separator))
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
// 整数 N が与えられるので、次の規則に従って N 行の出力をしてください。
// ・ N 行のうち、 i 行目では、1 から i までの数字を半角スペース区切りで出力してください。
// 例として、 N = 3 のとき、出力は次の通りになります。

// Input
// 3

// Output
// 1
// 1 2
// 1 2 3
