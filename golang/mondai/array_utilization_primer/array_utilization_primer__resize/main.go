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
	resArraySize := infoArray[1]

	var array []int
	for i := 0; i < max; i++ {
		str := nextLine()
		n, _ := strconv.Atoi(str)
		array = append(array, n)
	}

	var resArray []int
	for i := 0; i < resArraySize; i++ {
		if i < len(array) {
			resArray = append(resArray, array[i])
		} else {
			resArray = append(resArray, 0)
		}
	}

	for _, v := range resArray {
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
// 配列 A の要素数 N と新たに作成する配列のサイズ n ,
// 配列 A の各要素 A_1 ... A_N が与えられるので、
// 配列 A の先頭から n 要素を順に保持する配列を作成してください。
// 新たに作成する配列の要素数が A の要素数よりも大きい時は、
// サイズが合うように 0 を A の要素の末尾に追加してください。

// Input1
// 1 19
// 1

// Output1
// 1
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0
// 0

// Input2
// 5 3
// 1
// 2
// 3
// 4
// 5

// Output2
// 1
// 2
// 3
