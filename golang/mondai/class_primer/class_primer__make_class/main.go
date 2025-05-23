package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Employee struct {
	Number int
	Name   string
}

func main() {
	// 入力を受け取る
	commandCount := nextLineWithoutEmptySpace("int").([]int)[0]

	employees := map[int]Employee{}
	employeeCount := 0
	for i := 0; i < commandCount; i++ {
		command := nextLineWithoutEmptySpace("string").([]string)
		if command[0] == "make" {
			employeeCount++
			number, _ := strconv.Atoi(command[1])
			name := command[2]
			employees[employeeCount] = Employee{Number: number, Name: name}
		} else if command[0] == "getnum" {
			registerID := atoi(command[1])
			fmt.Println(employees[registerID].Number)
		} else if command[0] == "getname" {
			registerID := atoi(command[1])
			fmt.Println(employees[registerID].Name)
		}
	}
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}

	return i
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 空白を除いた行を取得してinterface{}で返却
func nextLineWithoutEmptySpace(elementType string) interface{} {
	line := nextLine()

	var numberArray []string
	numberArray = strings.Fields(line) // 空白を除去して分割

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
// エンジニアであり、社員を管理を管理する立場にあるあなたは、効率的に社員を管理するために、
// 各社員の社員番号 number と名前 name を持ち、加えて情報を返す関数を持つような構造体、すなわち次のようなメンバ変数とメンバ関数を持つ社員クラス class employee を作成することにしました。

// メンバ変数

// number : 整数
// name : 文字列

// メンバ関数
// getnum(){
//     return number;
// }
// getname(){
//     return name;
// }

// 入力で make number name と入力された場合は変数に number , name を持つ社員を作成し、 getnum n と入力された場合は n 番目に作成された社員の number を、getname n と入力された場合は n 番目に作成された社員の name を出力してください。

// Input1
// 7
// make 2742 mako
// getnum 1
// make 2782 taisei
// getname 2
// make 31 megumi
// getname 1
// getname 3

// Output1
// 2742
// taisei
// mako
// megumi
