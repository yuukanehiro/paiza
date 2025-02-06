package main

import (
	"fmt"
)

func main() {
	// 1文字ずつデータ型を指定して受け取る
	var a, b, c int                   // int型の変数を宣言
	var s string                      // string型の変数を宣言
	fmt.Scanf("%d %d %d", &a, &b, &c) // %dでint型を代入
	fmt.Scanf("%s", &s)               // %sでstring型を代入
}
