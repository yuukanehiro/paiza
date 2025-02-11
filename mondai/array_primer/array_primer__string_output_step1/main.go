package main

import (
	"fmt"
)

func main() {
	array := []string{
		"Hello",
		"paiza",
		"1234",
		"pa13",
	}

	fmt.Printf("%d\n", len(array))
}

// Q
// 複数の文字列があります。文字列の数を出力してください。
// Hello
// paiza
// 1234
// pa13

// Output:
// 4