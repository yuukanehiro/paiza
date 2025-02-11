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
	const separator = " "
	str := "1 3 5 1 2 3 6 6 5 1 4"

	array := strings.Split(str, separator)

	arrayInt := []int{}
	for _, v := range array {
		n, _ := strconv.Atoi(v)
		arrayInt = append(arrayInt, n)
	}

	uniqueArray := removeDuplicate(arrayInt)

	sortAsc(uniqueArray)

	for _, v := range uniqueArray {
		fmt.Printf("%d\n", v)
	}
}

// 昇順ソート
func sortAsc(numbers []int) {
	sort.Ints(numbers)
}

// 重複を削除してSliceを返却
func removeDuplicate[T comparable](array []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, v := range array {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
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
// 以下のような数列があります。

// 1 3 5 1 2 3 6 6 5 1 4

// この数列から数の重複をなくし、昇順にし改行区切りで出力してください。
// 数列を配列に格納し、並び替える操作や重複を削除する操作を考えて解いてみましょう。

// Output
// 1
// 2
// 3
// 4
// 5
// 6
