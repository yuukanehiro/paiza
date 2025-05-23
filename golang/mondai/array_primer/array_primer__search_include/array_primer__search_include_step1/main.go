package main

import (
	"fmt"
	"strings"
)

func main() {
	const separator = " "
	const targetNumber = "6"
	const answerYes = "Yes"
	const answerNo = "No"

	line := "10 13 21 1 6 51 10 8 15 6"
	array := strings.Split(line, separator)

	for _, v := range array {
		if v == targetNumber {
			fmt.Printf("%s\n", answerYes)
			return
		}
	}
}

// Q
// 以下のような配列があります。
// 10 13 21 1 6 51 10 8 15 6
// この中に、6 が含まれているなら Yes、含まれていないなら No を出力してください。

// Output
// Yes
