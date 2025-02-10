package main

// 配列を逆順にする
func reverse(array []int) {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-i-1] = array[len(array)-i-1], array[i]
	}
}

// target番目にvalueを挿入
func insertElement(array []int, target int, value int) []int {
	return append(array[:target-1], append([]int{value}, array[target-1:]...)...)
}

// target番目の要素を削除
func eraseElement(array []int, target int) []int {
	return append(array[:target-1], array[target:]...)
}
