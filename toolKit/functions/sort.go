package main

import "sort"

// 昇順ソート
func sortAsc(numbers []int) {
	sort.Ints(numbers)
}

// 降順ソート
func sortDesc(numbers []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
}
