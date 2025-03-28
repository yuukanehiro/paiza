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
	_ = nextLine()

	const separator = " "
	const elementType = "int"
	infoArray := nextLineBySeparator(separator, elementType).([]int)

	n := len(infoArray)

	// 行列の初期化
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		matrix[0][i] = infoArray[i]
		matrix[i][0] = infoArray[i]
	}

	if len(matrix) == 0 {
		fmt.Println("matrix is empty")
		return
	}

	// 結果の行列の初期化
	resultMatrix := make([][]string, n)
	for i := 0; i < n; i++ {
		resultMatrix[i] = make([]string, n)
	}

	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			resultMatrix[j][i] = strconv.Itoa(matrix[j][0]*matrix[0][i])
		}
	}

	for _, row := range resultMatrix {
		fmt.Println(strings.Join(row, separator))
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
