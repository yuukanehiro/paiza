package main

import "fmt"

func main() {
	array := []int{12, 3, 17, 9, 20, 6, 14, 1, 8, 15, 19, 4, 11, 5, 18, 13, 7, 2, 10, 16}

	array = quickSort(array)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func quickSort(array []int) []int {

	if len(array) < 2 {
		return array
	}

	pivot := array[len(array)/2]

	var less []int
	var equal []int
	var greater []int

	for _, v := range array {
		if v < pivot {
			less = append(less, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else if v > pivot {
			greater = append(greater, v)
		}
	}

	return append(append(quickSort(less), equal...), quickSort(greater)...)
}

// Output:
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// 11
// 12
// 13
// 14
// 15
// 16
// 17
// 18
// 19
// 20
