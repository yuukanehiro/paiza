package main

// 絶対値を返却
// 例) 10 -> 10, -10 -> 10
func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
