package main

// 配列の中に指定した要素が何個あるかをintで返却
func countTargetNumber[T comparable](array []T, targetNumber T) int {
	var targetCount int

	for _, v := range array {
		if targetNumber == v {
			targetCount++
		}
	}

	return targetCount
}

// 配列の全ての要素の和を返却
func sum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}

	return sum
}

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
		if v < min {
			min = v
		}
	}

	return min
}

// 配列の平均値を取得
func getAverage(array []int) float64 {
	if len(array) == 0 {
		return 0.0
	}

	sum := 0
	for _, v := range array {
		sum += v
	}

	return float64(sum) / float64(len(array))
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

// 配列の要素の種類数を返却
// 例: getKindCount([]int{1, 2, 3, 2, 1}) => 3
func getKindCount(array []int) int {
	seen := make(map[int]bool)

	for _, v := range array {
		if !seen[v] {
			seen[v] = true
		}
	}

	return len(seen)
}
