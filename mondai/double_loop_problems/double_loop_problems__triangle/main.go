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
	const separator = " "
	const elementType = "int"
	const triangleExist = "YES"
	const triangleNotExist = "NO"


	infoArray := nextLineBySeparator(separator, elementType).([]int)
	targetNum := infoArray[0]

	for i := 1; i <= targetNum; i++ {
		for j := 1; j <= targetNum; j++ {
			if i+j > targetNum {
				break
			}

			k := targetNum - i - j
			if k <= 0 {
				continue
			}

			if i * i == j * j + k * k || j * j == i * i + k * k || k * k == i * i + j * j {
				fmt.Println(triangleExist)
				return
			}
		}
	}

	fmt.Println(triangleNotExist)
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
// 整数 N が与えられるので、三角形の三辺の長さの和が N であり、
// 全ての辺の長さが整数であるような直角三角形が存在するかどうかを判定してください。
// なお、直角三角形の斜辺 a と他の二辺 b , c の間には次のような三平方の定理が成り立ちます。
// a ^ 2 = b ^ 2 + c ^ 2
// ・ ヒント
// 三辺の長さの和が N であるような全ての三角形の三辺 a , b , c の組み合わせのうち、
// 三平方の定理を満たすものが 1 つでもあれば "YES" ,
// それ以外の場合は "NO" が答えとなります。
// 全ての三辺の場合を全列挙することができれば三平方の定理を満たすかの判定をすることで答えを求めることができます。
