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
	coordinateMaxCount := infoArray[0]

	coordinateACount := infoArray[1]
	coordinateAIndex := coordinateACount - 1

	coordinateBCount := infoArray[2]
	coordinateBIndex := coordinateBCount - 1

	var coordinateArray []string
	for i := 0; i < coordinateMaxCount; i++ {
		tmp := nextLine()
		coordinateArray = append(coordinateArray, tmp)
	}

	coordinateAArray := strings.Split(coordinateArray[coordinateAIndex], " ")
	pointA_x, _ := strconv.Atoi(coordinateAArray[0])
	pointA_y, _ := strconv.Atoi(coordinateAArray[1])
	coordinateBarray := strings.Split(coordinateArray[coordinateBIndex], " ")
	pointB_x, _ := strconv.Atoi(coordinateBarray[0])
	pointB_y, _ := strconv.Atoi(coordinateBarray[1])

	fmt.Printf("%d\n", getManhattanDistance([]int{pointA_x, pointA_y}, []int{pointB_x, pointB_y}))
}

// マンハッタン距離を取得
// 解説: 2点間のx座標の差の絶対値とy座標の差の絶対値を足したもの
// 例) (1, 1), (2, 2) -> 2
func getManhattanDistance(pointA []int, pointB []int) int {
	return abs(pointA[0]-pointB[0]) + abs(pointA[1]-pointB[1])
}

// 絶対値を返却
// 例) 10 -> 10, -10 -> 10
func abs(n int) int {
	if n < 0 {
		return -n
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
// 1 行目に整数 N, A, B が与えられます。
// 2 行目以降に N 個の点の座標 x_1 y_1, x_2 y_2, ..., x_N y_N が与えられます。
// A 番目の点 と B 番目の点の距離を出力してください。

// 距離の計算にはマンハッタン距離

// |x1 - x2| + |y1 - y2|

// を用いることとします。

// 入力例1
// 3 1 3
// 2 3
// 1 2
// 5 6

// 出力例1
// 6
