package main

// 配列の中に指定した要素が含まれているかをboolで返却
func sliceContains[T comparable](slice []T, key T) bool {
	for _, s := range slice {
		if s == key {
			return true
		}
	}
	return false
}

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
