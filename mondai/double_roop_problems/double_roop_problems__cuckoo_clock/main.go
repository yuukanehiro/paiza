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
	for i := 0; i < 24; i++ {
		for j := 0; j < 60; j++ {
			fmt.Println(cuckooClock(i, j))
		}
	}
}

func cuckooClock(hour int, minute int) string {
	n := hour + minute

	if n%3 == 0 && n%5 == 0 {
		return "FIZZBUZZ"
	} else if n%3 == 0 {
		return "FIZZ"
	} else if n%5 == 0 {
		return "BUZZ"
	}

	return ""
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
// 普通の鳩時計は 1 時間に 1 回しか鳴かないのでつまらないと思ったあなたは、
// 鳩時計を改造してスーパー鳩時計を作りました。
// このスーパー鳩時計は時刻が x 時 y 分のとき x + y が 3の倍数のとき"FIZZ"、
// 5 の倍数のとき"BUZZ", 
// 3の倍数かつ5の倍数のとき "FIZZBUZZ" と鳴き、
// これらのいずれにも当てはまらなかった場合は鳴きません。
// なお、0 は 3 の倍数かつ 5 の倍数であるとします。 
// 0 時 0 分　〜 23 時 59 分 の各分のスーパー鳩時計の様子を出力してください。
