package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type User struct {
	nickname string
	old      int
	birth    string
	state    string
}

func main() {
	infoArray := nextLineBySeparator(" ", "int").([]int)
	rowCount := infoArray[0]
	changeRowCount := infoArray[1]

	var userList []User
	for i := 0; i < rowCount; i++ {
		userArray := nextLineBySeparator(" ", "string").([]string)
		user := User{
			nickname: userArray[0],
			old: func(old string) int {
				i, e := strconv.Atoi(old)
				if e != nil {
					panic(e)
				}
				return i
			}(userArray[1]),
			birth: userArray[2],
			state: userArray[3],
		}
		userList = append(userList, user)
	}

	changeList := make([][]string, changeRowCount)
	for i := 0; i < changeRowCount; i++ {
		changeArray := nextLineBySeparator(" ", "string").([]string)
		// 初期化
		changeList[i] = []string{}
		changeList[i] = changeArray
	}

	changeUserName(userList, changeList)

	for _, user := range userList {
		fmt.Printf("%s %d %s %s\n", user.nickname, user.old, user.birth, user.state)
	}
}

func changeUserName(userList []User, changeList [][]string) {
	for _, change := range changeList {
		changeRowCount := change[0]
		changeRowCountInt, err := strconv.Atoi(changeRowCount)
		if err != nil {
			panic(err)
		}
		// 行数をIndexに変換
		changeRowCountInt--
		changeUserName := change[1]
		userList[changeRowCountInt].nickname = changeUserName
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

// 送ってもらったデータを使いやすいように整理したいと思った paiza 君はクラス全員分のデータを次の形式でまとめることにしました。

// User{
//     nickname : 名前
//     old : 年齢
//     birth : 誕生日
//     state : 出身地
// }

// 途中で名前が変わった際にいちいちデータを修正するのが面倒だと思ったあなたは、生徒の構造体と新しい名前を受け取り、その名前を修正する関数 changeName を作成し、それを用いて生徒の名前を更新することにしました。

// クラスの人数と全員の情報、更新についての情報が与えられるので、入力に従って名前を更新した後のクラスのメンバーの情報を出力してください。

// Input1
// 1 1
// koko 23 04/10 tokyo
// 1 nana

// Output1
// nana 23 04/10 tokyo

// Input2
// 3 2
// mako 13 08/08 nara
// taisei 16 12/04 nagano
// megumi 14 11/02 saitama
// 2 taihei
// 3 megu

// Output2
// mako 13 08/08 nara
// taihei 16 12/04 nagano
// megu 14 11/02 saitama
