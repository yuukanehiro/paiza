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
	array := nextLineBySeparator(" ", "int").([]int)
	max := array[0]

	var queue []int
	for i := 0; i < max; i++ {
		str := nextLine()
		if strings.Contains(str, "in") {
			n, _ := strconv.Atoi(strings.Split(str, " ")[1])
			queue = append(queue, n)
			continue
		}
		if strings.Contains(str, "out") && len(queue) > 0 {
			queue = queue[1:]
		}
	}

	for _, v := range queue {
		fmt.Printf("%d\n", v)
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

// Q
// データ構造の queue と同様の働きをするロボットがあります。
// ロボットは指示に応じて配列 A に対して 2 種類の仕事を行います、仕事の内容は以下の通りです。

// ・in n と指示された場合、A の末尾に n を追加してください。
// ・out と指示された場合、A の先頭の要素を削除してください。ただし、A が既に空の場合、何も行わないでください。

// ロボットに与えられる指示の回数 N と、各指示の内容 S_i が与えられるので、
// ロボットが全ての処理を順に行った後の A の各要素を出力してください。
// なお、初め配列 A は空であるものとします。

// Input
// 28
// in 43
// in -14
// in 9
// in 42
// out
// in 78
// out
// in -71
// in -26
// out
// out
// out
// in -22
// out
// in 47
// in -86
// out
// out
// out
// out
// in -26
// out
// out
// out
// in 81
// in -9
// out
// in -18

// Output
// -9
// -18
