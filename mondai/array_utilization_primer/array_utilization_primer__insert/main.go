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
	max := infoArray[0]

	var array []int
	for i := 0; i < max; i++ {
		str := nextLine()
		n, _ := strconv.Atoi(str)
		array = append(array, n)
	}

	targetInfoArray := nextLineBySeparator(" ", "int").([]int)
	target := targetInfoArray[0]
	targetNumber := targetInfoArray[1]

	array = insertElement(array, target, targetNumber)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

// target番目にvalueを挿入
func insertElement(array []int, target int, value int) []int {
	// target の範囲チェック
	if target < 0 || target > len(array) {
		fmt.Println("Error: target index out of range")
		return array // そのまま返す
	}

	return append(array[:target], append([]int{value}, array[target:]...)...)
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
// 配列 A と追加する位置 n と追加する要素 B が与えられるので、B を A_n の後ろに追加した後の A を出力してください。

// Input
// 1
// 1
// 1 2

// Output
// 1
// 2
