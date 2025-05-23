package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var userList []User
	for i := 0; i < rowCount; i++ {
		userArray := nextLineBySeparator(" ", "string").([]string)
		user := User{
			nickname: userArray[0],
			old:      atoi(userArray[1]),
			birth:    userArray[2],
			state:    userArray[3],
		}
		userList = append(userList, user)
	}

	// 年齢で昇順でソート
	sort.Slice(userList, func(i, j int) bool {
		return userList[i].old < userList[j].old
	})

	for _, user := range userList {
		fmt.Printf("%s %d %s %s\n", user.nickname, user.old, user.birth, user.state)
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
