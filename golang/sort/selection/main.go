package main

import "fmt"

func main() {
	array := []int{12, 3, 17, 9, 20, 6, 14, 1, 8, 15, 19, 4, 11, 5, 18, 13, 7, 2, 10, 16}

	array = selectionSort(array)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}

	return array
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
