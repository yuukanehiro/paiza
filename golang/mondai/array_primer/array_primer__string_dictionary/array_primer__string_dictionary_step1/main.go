package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	separator := " "
	arrayStr := "zaipa izapa paiza"
	array := strings.Split(arrayStr, separator)

	sortStrAsc(array)

	for _, v := range array {
		fmt.Printf("%s\n", v)
	}
}

// 辞書順にソート
func sortStrAsc(strings []string) {
	sort.Strings(strings)
}

// Q
// 3 つの文字列

// zaipa izapa paiza

// があります。辞書順に並べ替え、改行区切りで出力してください。
// この三つの文字列を配列に格納し、その配列を並び替える操作を考えて解いてみましょう。

// Output:
// izapa
// paiza
// zaipa