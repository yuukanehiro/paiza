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
	max := infoArray[0]

	type student struct {
		name string
		old int
		birth string
		state string
	}

	students := make([]student, max, max)
	for i := 0; i < max; i++ {
		studentInfo := nextLineBySeparator(" ", "string").([]string)
		students = append(students, student{
			name: studentInfo[0],
			old: func() int {
				old, _ := strconv.Atoi(studentInfo[1])
				return old
			}(),
			birth: studentInfo[2],
			state: studentInfo[3],
		})
	}

	targetOld := nextLine()
	targetOldInt, _ := strconv.Atoi(targetOld)

	for _, student := range students {
		if student.old == targetOldInt {
			fmt.Println(student.name)
		}
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
// クラスの学級委員である paiza 君は、クラスのみんなに次のような形式でアカウントの情報を送ってもらうよう依頼しました。

// 名前 年齢 誕生日 出身地

// 送ってもらったデータを使いやすいように整理したいと思った paiza 君はクラス全員分のデータを次のような構造体でまとめることにしました。

// student{
//     name : 名前
//     old : 年齢
//     birth : 誕生日
//     state : 出身地
// }


// 年齢ごとの生徒の名簿を作る仕事を任された paiza 君はクラスメイトのうち、決まった年齢の生徒を取り出したいと考えました。
// 取り出したい生徒の年齢が与えられるので、その年齢の生徒の名前を出力してください。
