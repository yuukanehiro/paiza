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
	arrMax := infoArray[0]
	commandMax := infoArray[1]

	var array []int
	for i := 0; i < arrMax; i++ {
		array = append(array, i + 1)
	}

	if len(array) == 0 {
		fmt.Println("Error: array length is less than 1")
		return
	}

	var commandArray []string
	for i := 0; i < commandMax; i++ {
		commandArray = append(commandArray, nextLine())
	}

	if len(commandArray) == 0 {
		fmt.Println("Error: commandArray length is less than 1")
		return
	}

	for _, command := range commandArray {
		if strings.Contains(command, "swap") {
			// "swap A B"
			commandSwapArray := strings.Split(command, " ")

			swapIndex1, _ := strconv.Atoi(commandSwapArray[1])
			swapIndex1--
			swapIndex2, _ := strconv.Atoi(commandSwapArray[2])
			swapIndex2--

			swap(array, swapIndex1, swapIndex2)
		} else if strings.Contains(command, "reverse") {
			reverseArray(array)
		} else if strings.Contains(command, "resize") {
			commandResizeArray := strings.Split(command, " ")
			resizeNumber, _ := strconv.Atoi(commandResizeArray[1])
			// 配列の長さが resizeNumber より大きい場合のみ resize
			if len(array) > resizeNumber {
				array = array[:resizeNumber]
			}
		}
	}

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func swap(array []int, index1 int, index2 int) {
	if index1 >= 0 && index1 < len(array) && index2 >= 0 && index2 < len(array) {
		array[index1], array[index2] = array[index2], array[index1]
	}
}

func reverseArray(array []int) {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
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
// あなたは集団行動のリーダーです。次のような指示を出すことで様々な列の操作ができます。

// ・swap A B
// 先頭から A 番目の人と、先頭から B 番目の人の位置を入れ替える。
// ・reverse
// 列の前後を入れ替える。
// ・resize C
// 先頭から C 人を列に残し、それ以外の人を全員列から離れさせる。
// ただし、列が既に C 人以下の場合、何も行わない。

// 初め、列には番号 1 〜 N の N 人がおり、
// 先頭から番号の昇順に並んでいます。(1, 2 , 3, ..., N)
// あなたの出した指示の回数 Q とその指示の内容 S_i (1 ≦ i ≦ Q) が順に与えられるので、
// 全ての操作を順に行った後の列を出力してください。

// Input
// 10 2
// reverse
// resize 7

// Output
// 10
// 9
// 8
// 7
// 6
// 5
// 4
