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

func NewEmployee(number int, name string) Employee {
	return Employee{
		Number: number,
		Name:   name,
	}
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
			number := atoi(command[1])
			name := command[2]
			employees[employeeCount] = NewEmployee(number, name)
		} else if command[0] == "getnum" {
			registerID := atoi(command[1])
			fmt.Println(employees[registerID].Number)
		} else if command[0] == "getname" {
			registerID := atoi(command[1])
			fmt.Println(employees[registerID].Name)
		} else if command[0] == "change_num" {
			registerID := atoi(command[1])
			NewNumber := atoi(command[2])
			employees[registerID] = NewEmployee(NewNumber, employees[registerID].Name)
		} else if command[0] == "change_name" {
			registerID := atoi(command[1])
			NewName := command[2]
			employees[registerID] = NewEmployee(employees[registerID].Number, NewName)
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
// エンジニアであり、社員を管理する立場にあるあなたが勤める会社には、効率的に社員を管理するために、次のようなメンバ変数とメンバ関数を持つ社員クラス class employee が存在しています。

// メンバ変数
// 整数 number, 文字列 name

// メンバ関数

// getnum(){
//     return number;
// }
// getname(){
//     return name;
// }

// しかし、この社員クラスについて、社員番号や名前が変わった際にいちいち手動で更新するのは面倒だと感じたあなたは、引数に元の社員番号と新しい社員番号を指定すれば、新しい社員番号に更新してくれる関数 change_num と 引数に元の名前と新しい名前を指定すれば、新しい名前に更新してくれる関数 change_name を作成することにしました。

// 入力で make number name と入力された場合は変数にnumber, nameを持つ社員を作成し、getnum nと入力された場合は n 番目に作成された社員の number を、getname n と入力された場合は n 番目に作成された社員の name を出力してください。

// また、 change_num n newnum と入力された場合は、n 番目に作成された社員の number を、 newnum に変更し、 change_name n newname と入力された場合は、n 番目に作成された社員の name を、 newname に変更してください。

// Input
// 12
// make 2742 makoto
// getnum 1
// make 2782 taro
// getname 1
// getname 2
// change_num 2 9927
// change_name 1 mako
// getnum 2
// make 31 meu
// change_name 3 meumeu
// getnum 3
// getname 1

// Output
// 2742
// makoto
// taro
// 9927
// 31
// mako
