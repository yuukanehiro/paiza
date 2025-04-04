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

// 辞書順にソート
func sortStrAsc(strings []string) {
	sort.Strings(strings)
}

// 構造体のソート
// 使用例
// 年齢で昇順でソート
//
//	type User struct {
//		nickname string
//		old      int
//	}
//
//	userList := []User{
//		{"Alice", 25},
//		{"Bob", 20},
//	}
//
//	sortSliceStructByFunc(userList, func(i, j int) bool {
//		return userList[i].old < userList[j].old
//	})
func sortSliceStructByFunc[T any](s []T, f func(i, j int) bool) {
	sort.Slice(s, f)
}
