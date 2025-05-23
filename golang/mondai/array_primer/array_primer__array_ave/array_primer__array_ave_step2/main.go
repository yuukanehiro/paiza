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
	targetArray := nextLineBySeparator(" ", "int").([]int)
	pivotNumber := targetArray[1]

	array := nextLineBySeparator(" ", "int").([]int)

	result := []int{}

	for _, v := range array {
		if v >= pivotNumber {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return
	}

	for _, v := range result {
		fmt.Printf("%d\n", v)
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
// 1 行目に整数 N, K が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、K 以上の数をすべて、入力された順に改行区切りで出力してください。
// また、K 以上の数が一個もない場合は、何も出力しなくても問題ありません。

// Input
// 5 3
// 5 2 4 3 1

// Output
// 5
// 4
// 3