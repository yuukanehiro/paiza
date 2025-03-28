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
	coin1 := 1
	coinX := infoArray[0]
	coinY := infoArray[1]
	price := infoArray[2]

	// coinXとcoinYがpriceよりも大きい場合はcoin1で支払う
	if coinX > price && coinY > price {
		fmt.Println(price / coin1)
		return
	}

	// 出題の条件外の場合はエラー
	if coinX == coinY {
		panic("coinX and coinY are the same")
	}

	// coin1, coinX, coinYでpriceを支払う場合の最小数枚数を求める
	var minCount int = price/coin1
	for i := 0; i <= price/coin1; i++ {
		for j := 0; j <= price/coinX; j++ {
			for k := 0; k <= price/coinY; k++ {
				if coin1*i+coinX*j+coinY*k == price {
					count := i + j + k
					if count < minCount {
						minCount = count
					}
				}
			}
		}
	}

	fmt.Println(minCount)
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
// paiza国では、1 円と X 円と Y 円の 3 種類の硬貨しかありません。
// ちょうど Z 円を支払うとき、
// 支払う硬貨の枚数が最小になるように支払ったときの硬貨の枚数を求めてください。
// ただし、支払う各硬貨の枚数に制限は無いものとします。
