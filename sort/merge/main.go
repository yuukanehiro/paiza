package main

import "fmt"

func main() {
	array := []int{12, 3, 17, 9, 20, 6, 14, 1, 8, 15, 19, 4, 11, 5, 18, 13, 7, 2, 10, 16}

	array = mergeSort(array)

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2

	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// 左右の配列を比較しながらマージ
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 残りの要素を追加
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
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
