package main

// マンハッタン距離を取得
// 解説: 2点間のx座標の差の絶対値とy座標の差の絶対値を足したもの
// 例) (1, 1), (2, 2) -> 2
func getManhattanDistance(pointA []int, pointB []int) int {
	return abs(pointA[0]-pointB[0]) + abs(pointA[1]-pointB[1])
}
