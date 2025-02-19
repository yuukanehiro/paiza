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
	infoArray1 := nextLineBySeparator(" ", "int").([]int)
	arrayMax := infoArray1[0]
	examWeightingArray := nextLineBySeparator(" ", "int").([]int)

	var studentExamResultArray [][]int
	for i := 0; i < arrayMax; i++ {
		studentExamResultArray = append(studentExamResultArray, nextLineBySeparator(" ", "int").([]int))
	}

	var kingResult int = 0
	for _, v := range studentExamResultArray {
		var total int
		for i, vv := range v {
			total += vv * examWeightingArray[i]
		}

		if total > kingResult {
			kingResult = total
		}
	}

	fmt.Printf("%d\n", kingResult)
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
// paiza の入社試験では 科目 1 〜 5 の 5 科目のテストが課せられており、
// それぞれの科目には重みが設定されています。受験者の得点は各科目の (とった点数) * (科目の重み) となります。
// 5 科目の得点の合計が最も高かった受験者の得点を求めてください。

// Input1
// 5
// 1 2 3 4 5
// 1 2 3 4 5
// 6 7 8 9 0
// 10 11 12 13 14
// 10 10 10 10 10
// 100 2 4 6 48

// Output1
// 380

// Input2
// 1
// 1 2 3 4 5
// 100 100 100 100 100

// Output2
// 1500
