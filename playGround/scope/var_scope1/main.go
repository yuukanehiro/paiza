package main

import (
	"fmt"
)

func main() {
	var array []int
	array = []int{1, 2, 3}

	trueFlag := true

	if trueFlag {
		array := []int{9}
		_ = array
	}

	for _, v := range array {
		fmt.Println(v)
	}
}

// Output
// 1
// 2
// 3

// array := []int{9} で array が再定義されているため、
// for文で出力されるのは、再定義前の array である [1, 2, 3] となる。
// そのため、出力は以下の通り。
// 1
// 2
// 3
