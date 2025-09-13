package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s) // [1 2 3 4 5 6 7 8 9 10]

	fmt.Println(s[0])        // 1
	fmt.Println(s[1])        // 2
	fmt.Println(s[len(s)-1]) // 10
	fmt.Println(s[:3])       // [1 2 3]
	fmt.Println(s[3:6])      // [4 5 6]
	fmt.Println(s[6:])       // [7 8 9 10]
}
