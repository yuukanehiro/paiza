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
	pivot := infoArray[1]

	var array []int
	for i := 0; i < max; i++ {
		str := nextLine()
		n, _ := strconv.Atoi(str)
		array = append(array, n)
	}

	var filteredArray []int
	for _, v := range array {
		if v >= pivot {
			filteredArray = append(filteredArray, v)
		}
	}

	for _, v := range filteredArray {
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
// 配列 A の要素数 N と整数 K , 配列 A の各要素 A_1 ... A_N が与えられるので、
// K 以上であるような配列 A の要素のみを要素として持つ配列 B を作成し、
// その要素を出力してください。 B の要素の順は A と同じにしてください。

// Input
// 5 3
// 1
// 2
// 3
// 4
// 5

// Output
// 3
// 4
// 5
