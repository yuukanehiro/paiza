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
	Nickname string
	Old int
	Birth string
	State string
}

func (u User) String() string {
	return fmt.Sprintf(`User{
nickname : %s
old : %d
birth : %s
state : %s
}`, u.Nickname, u.Old, u.Birth, u.State)}

func main() {
	infoArray := nextLineBySeparator(" ", "int").([]int)
	max := infoArray[0]

	// 初期化
	var userInfoArray [][]string
	for i := 0; i < max; i++ {
		userInfoArray = append(userInfoArray, []string{})
		userInfoArray[i] = nextLineBySeparator(" ", "string").([]string)
	}

	var users []User
	for i := 0; i < max; i++ {

		users = append(users, User{
			Nickname: userInfoArray[i][0],
			Old: func() int {
				old, _ := strconv.Atoi(userInfoArray[i][1])
				return old
			}(),
			Birth: userInfoArray[i][2],
			State: userInfoArray[i][3],
		})
	}

	for _, user := range users {
		fmt.Println(user.String())
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
// nickname : 名前
// old : 年齢
// birth : 誕生日
// state : 出身地
// }


// クラスメートの情報が与えられるので、それらを以上の形式でまとめたものを出力してください。
