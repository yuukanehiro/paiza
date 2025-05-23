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
	limit := infoArray[0]

	var array []int
	for i := 0; i < limit; i++ {
		str := nextLine()
		n, _ := strconv.Atoi(str)

		array = append(array, n)
	}

	reverse(array)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

// 配列を逆順にする
func reverse(array []int) {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-i-1] = array[len(array)-i-1], array[i]
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
// 配列 A の要素数 N と配列 A の各要素 A_1, A_2, ..., A_N が与えられるので、
// A の要素の順序を逆にした配列 B を作成し、出力してください。

// Input
// 5
// 1
// 2
// 3
// 4
// 5

// Output
// 5
// 4
// 3
// 2
// 1
