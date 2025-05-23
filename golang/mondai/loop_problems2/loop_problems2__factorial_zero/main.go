package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

// 5の冪乗のリストを作成し、そのリストに含まれる数で割り切れる数をカウントする
// 5の冪乗のリストを作成する際に、targetより大きい数は作成しない
func main() {
	infoArray := nextLineBySeparator(" ", "int").([]int)
	target := infoArray[0]

	// 5の冪乗のリストを作成
	var fivePowers []int
	var fivePower int
	for i := 5; i <= target; i *= 5 {
		fivePower = i
		fivePowers = append(fivePowers, fivePower)
	}

	var count int
	for i := target; i > 0; i-- {
		for _, v := range fivePowers {
			if i%v == 0 {
				count++
			}
		}
	}

	fmt.Println(count)
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

// input1
// 100

// output1
// 24
// 内訳: 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100
// 5の冪乗のリスト: 5, 25
// 5で割り切れる数: 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100
// 25で割り切れる数: 25, 50, 75, 100
// 合計: 24
