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
	intArray := nextLineBySeparator(" ", "int").([]int)
	targetNumber := intArray[0]
	targetIndex := targetNumber - 1

	targetWordCount := intArray[2]

	strArray := nextLineBySeparator(" ", "string").([]string)
	targetWord := strArray[targetIndex]

	fmt.Printf("%s\n", getCharAt(targetWord, targetWordCount))
}

// i番目の文字を取得
// 例: getRuneAt("paiza", 4) => "z"
func getCharAt(s string, i int) string {
	rs := []rune(s)
	return string(rs[i-1])
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
// 1 行目に整数 N, M, L が与えられます。
// 2 行目に M 個の文字列 s_1, s_2, ..., s_M が半角スペース区切りで与えられます。
// N 番目の文字列 s_N の L 番目の文字を出力してください。

// Input
// 3 5 2
// abc def ghi jkl mno

// Output
// h