package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/big"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	infoArray := nextLineBySeparator(" ", "int").([]int)
	targetNum := infoArray[0]

	m := factorial(targetNum)

	// 2を定に持つ対数の計算
	fmt.Println(countDivisibleExponent(int64(2), m))
}

// 指定した指数の数を計算
// 例) 10!の中に2がいくつあるか
// 10! = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
// 10! = 2^8 * 3^4 * 5^2 * 7
// 2の数は8個
// 2, 10! -> 8
// int64(2), big.NewInt(3628800) -> 8
func countDivisibleExponent(exponent int64, n *big.Int) int {
	zero := big.NewInt(0)
	baseNum := big.NewInt(exponent)
	count := 0

	if n.Cmp(zero) == 0 {
		return 0
	}

	tmp := new(big.Int).Set(n)
	mod := new(big.Int)

	for {
		// 余りが0でない場合はループを抜ける
		mod.Mod(tmp, baseNum)
		if mod.Cmp(zero) != 0 {
			break
		}

		// 余りが0の場合は割る
		tmp.Div(tmp, baseNum)
		count++
	}

	return count
}

// 階乗の計算
// 例) 5 -> 120
// intで再起処理を行うと、nが大きい場合にスタックオーバーフローが発生するため、big.Intを利用
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
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
