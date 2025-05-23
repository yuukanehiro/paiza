package main

// i番目の文字を取得
// 例: getRuneAt("paiza", 4) => "z"
func getCharAt(s string, i int) string {
	rs := []rune(s)
	return string(rs[i-1])
}
