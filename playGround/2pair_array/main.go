package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	// 1番目を除いて2個ずつペアを作る
	pairs := make(map[int][]int, 0)
	for i := 0; 1+i*2 < len(array); i++ {
		pairs[i] = make([]int, 2)
		pairs[i][0] = array[1+i*2]
		pairs[i][1] = array[2+i*2]
	}

	b, _ := json.MarshalIndent(pairs, "", "  ")
	fmt.Println(string(b))

	// {
	// 	"0": [
	// 	  2,
	// 	  3
	// 	],
	// 	"1": [
	// 	  4,
	// 	  5
	// 	],
	// 	"2": [
	// 	  6,
	// 	  7
	// 	],
	// 	"3": [
	// 	  8,
	// 	  9
	// 	],
	// 	"4": [
	// 	  10,
	// 	  11
	// 	]
	//   }
}
