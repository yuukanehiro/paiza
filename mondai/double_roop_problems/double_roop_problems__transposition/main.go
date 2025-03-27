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
	rowCount := infoArray[0]

	var matrix [][]int
	for i := 0; i < rowCount; i++ {
		matrix = append(matrix, nextLineBySeparator(" ", "int").([]int))
	}

	if len(matrix) == 0 {
		fmt.Println("matrix is empty")
		return
	}

	var transposedMatrix [][]string
	for i := 0; i < len(matrix[0]); i++ {
		var row []string
		for j := 0; j < len(matrix); j++ {
			row = append(row, strconv.Itoa(matrix[j][i]))
		}
		transposedMatrix = append(transposedMatrix, row)
	}

	for _, row := range transposedMatrix {
		fmt.Println(strings.Join(row, " "))
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
