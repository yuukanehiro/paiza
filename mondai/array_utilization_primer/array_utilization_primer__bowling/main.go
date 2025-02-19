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
	array4 := nextLineBySeparator(" ", "int").([]int)
	reverseArray(array4)
	array3 := nextLineBySeparator(" ", "int").([]int)
	reverseArray(array3)
	array2 := nextLineBySeparator(" ", "int").([]int)
	reverseArray(array2)
	array1 := nextLineBySeparator(" ", "int").([]int)
	reverseArray(array1)


	var array []int
	array = append(array, append(array1, append(array2, append(array3, array4...)...)...)...)

	var firstPinCount int
	for i, v := range array {
		if v == 1 {
			firstPinCount = i + 1
			break
		}
	}

	enablePinCount := 0
	for _, v := range array {
		if v == 1 {
			enablePinCount++
		}
	}

	fmt.Printf("%d\n", firstPinCount)
	fmt.Printf("%d\n", enablePinCount)
}

// 配列を逆順にする
func reverseArray(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
	return array
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
// あなたはボウリングをしています。
// フレームの 1 投目を投げ終わったあなたは、
// 次に狙うピンの番号と残っているピンの本数を知りたくなりました。
// ピンの情報が与えられるので、狙うべきピンの番号と残っているピンの本数を求めてください。

// 狙うピンの決め方は次の通りとします。
// - 残っているピンのうち、先頭 (P_1側) のピンを狙います。
// ただし、同じ列に複数ピンがある場合は、それらのうちピン番号が最も小さいピンを狙います。

// 入力される値
// P_10 P_9 P_8 P_7
// P_6 P_5 P_4
// P_3 P_2
// P_1

// Input
// 1 0 0 1
// 0 0 0
// 0 0
// 0

// Output
// 7
// 2
