package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	intArray := nextLineBySeparator(" ", "int").([]int)
	targetNumber := intArray[1]
	targetIndex := targetNumber - 1

	array := nextLineBySeparator(" ", "string").([]string)
	sortStrAsc(array)

	fmt.Printf("%s\n", array[targetIndex])
}

// 辞書順にソート
func sortStrAsc(strings []string) {
	sort.Strings(strings)
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
// 2 行目に N 個の文字列 s_1, s_2, ..., s_N が半角スペース区切りで与えられます。
// N 個の文字列を辞書順に並べ替え、K 番目の文字列を出力してください。

// Input
// 5 2
// e d c b a

// Output
// b