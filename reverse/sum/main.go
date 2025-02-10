package main

import "fmt"

func main() {
	n := 5

	result := sum(n)

	fmt.Printf("%d\n", result)
}

func sum(n int) int {
	if n == 0 {
		return 0
	}

	return n + sum(n-1)
}

// Output:
// 15
