package main

// 配列の最大値を取得
func getMax(array []int) int {
	max := array[0]

	for _, v := range array {
		if v > max {
			max = v
		}
	}

	return max
}

// 配列の最小値を取得
func getMin(array []int) int {
	min := array[0]

	for _, v := range array {
		if v < min{
			min = v
		}
	}

	return min
}

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

// 重複を削除してSliceを返却
// 例: removeDuplicate([]int{1, 2, 3, 2, 1}) => []int{1, 2, 3}
func removeDuplicate[T comparable](array []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, v := range array {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}