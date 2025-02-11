package main

import (
	"fmt"
)

func main() {
	array := []string{
		"eight",
		"one",
		"three",
		"paiza",
		"pa13",
		"813",
	}

	for _, v := range array {
		fmt.Printf("%s\n", v)
	}
}

// Q
// 複数の文字列があります。すべての文字列を改行区切りで出力してください。

// eight
// one
// three
// paiza
// pa13
// 813

// Output:
// eight
// one
// three
// paiza
// pa13
// 813