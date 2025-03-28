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
	max := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i + j >= 100 || i*i*i + j*j*j >= 100000 {
				break
			}

			if i * j > max {
				max = i * j
			}
		}
	}

	fmt.Println(max)
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
// x + y < 100 かつ (x ^ 3) + (y ^ 3) < 100000 が成り立つような正の整数 x , y について
//  x × y の最大値を求めてください。

// ・ ヒント
// 2 つの式を連立不等式として解きたくなりますが、
// x + y < 100 に注目すると、(x , y) のとりうる値は 
// (1,1) , (1,2) , (1,98) , (2,1)... (98,1) のいずれかであり、
// これらは高々 98 + 97 + ... + 1 = 99 × 44 = 4356 通り（等差数列の和の公式を利用）であるため、
// 全ての組を調べても実行時間制限に間に合います。
