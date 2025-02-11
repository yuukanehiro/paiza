package main

import (
	"fmt"
)

func main() {
	array := []string{
		"good",
		"morning",
		"paiza",
		"813",
		"pa13",
	}

	for _, v := range array {
		fmt.Printf("%s\n", v)
	}
}

// Q
// 複数の文字列があります。すべての文字列を改行区切りで出力してください。

// good
// morning
// paiza
// 813
// pa13

// Output:
// good
// morning
// paiza
// 813
// pa13