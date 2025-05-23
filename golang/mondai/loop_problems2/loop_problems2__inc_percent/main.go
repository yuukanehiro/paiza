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
	const incrementPercent10 = 1.1
	infoArray := nextLineBySeparator(" ", "int").([]int)
	startYen := infoArray[0]
	limitYen := infoArray[1]

	var count int = 0
	for i := startYen; i <= limitYen; i = int(float64(i) * incrementPercent10) {
		count += 1
	}

	fmt.Printf("%d\n", count)
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
// 現在所持金を A 円持っています。
// 所持金が毎日 10% ずつ増えるとき、何日後に B 円を超えるか出力してください。
// また、増加するお金は小数点以下切り捨てで考えることとします。
// 例として、所持金が 831 円 のとき、10% は 83.1円 ですが、増加するお金は 83 円 です。

// Input
// 10 15

// Output
// 6
