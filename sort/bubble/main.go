package main

import "fmt"

func main() {
	var array = []int{10, 2, 3, 1, 7, 5, 6, 4, 9, 8}

	bubbleSort(array)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
