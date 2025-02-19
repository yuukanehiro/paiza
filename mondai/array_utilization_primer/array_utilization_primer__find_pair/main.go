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

	if len(array) < 2 {
		fmt.Println("Error: array length is less than 2")
		return
	}

	var minDifference int = 999
	var minPair []int
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			diff := abs(array[i] - array[j])
			if diff < minDifference {
				minDifference = diff
				if array[i] < array[j] {
					minPair = []int{array[i], array[j]}
				} else {
					minPair = []int{array[j], array[i]}
				}
			}
		}
	}

	for _, v := range minPair {
		fmt.Printf("%d\n", v)
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
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
// 生徒の身長が A_1, ...., A_N であるような N 人のクラスで二人三脚の代表を決めることにしました。
// 代表には、身長の差が最も小さいような 2 人を選出することにしました。
// 選ばれた 2 人の身長を昇順に出力してください。

// Input
// 5
// 119
// 102
// 187
// 191
// 132

// Output
// 187
// 191
