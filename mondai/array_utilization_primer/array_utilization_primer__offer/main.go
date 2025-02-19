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
	declinedCount := infoArray[2]

	var array []int
	for i := 0; i < max; i++ {
		str := nextLine()
		n, _ := strconv.Atoi(str)
		array = append(array, n)
	}

	passCount := 0
	for _, v := range array {
		if v >= pivot {
			passCount++
		}
	}

	recruitCount := passCount - declinedCount
	if recruitCount < 0 {
		recruitCount = 0
	}

	fmt.Println(recruitCount)
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
// 人事のあなたは、N 人の中から採用者を決定します。
// N 人のテストの点数はそれぞれ A_i (1 ≦ i ≦ N)です。
// テストの点数が K 点以上の人全員を採用したいのですが、得点の高い方から M 人に辞退されてしまいました。
// あなたの採用することのできる最大の人数を答えてください。
// 採用できる人数が 0 人の場合もあることに気をつけてください。

// Input
// 5 3974 0
// 2049
// 4690
// 6867
// 3414
// 460

// Output
// 2
