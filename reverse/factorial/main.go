package main

import "fmt"

func main() {
	num := 5 // 5! = 120

	result := factorial(num)

	fmt.Printf("%d\n", result)
}

// 階乗を計算して返却
// 5! = 5 * 4 * 3 * 2 * 1 = 120
// 5! = 5 * 4!
func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

// Output:
// 120
