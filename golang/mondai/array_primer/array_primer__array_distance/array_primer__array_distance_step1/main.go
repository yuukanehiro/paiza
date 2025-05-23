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
	pointCount := infoArray[0]

	var coordinateArray []string
	for i := 0; i < pointCount; i++ {
		tmp := nextLine()
		coordinateArray = append(coordinateArray, tmp)
	}

	pivot_x := 2
	pivot_y := 3

	var distanceArray []int
	for i := 0; i < pointCount; i++ {
		coordinate := strings.Split(coordinateArray[i], " ")
		x, _ := strconv.Atoi(coordinate[0])
		y, _ := strconv.Atoi(coordinate[1])

		distance := abs(pivot_x-x) + abs(pivot_y-y)
		distanceArray = append(distanceArray, distance)
	}

	for _, v := range distanceArray {
		fmt.Printf("%d\n", v)
	}
}

// 絶対値を返却
func abs(n int) int {
	if n < 0 {
		return -n // 例) -(-10) = 10
	}

	return n
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
// 1 行目に整数 N が与えられます。
// 2 行目以降に N 個の点の座標 x_1 y_1, x_2 y_2, ..., x_N y_N が与えられます。
// 点 (2, 3) と各点の距離を改行区切りで出力してください。

// 距離の計算にはマンハッタン距離

// |x1 - x2| + |y1 - y2|

// を用いることとします。

// Input
// 3
// 2 3
// 1 2
// 5 6

// Output
// 0
// 2
// 6
