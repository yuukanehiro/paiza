package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	infoArray := nextLineInts()
	max := infoArray[0]
	target := infoArray[1]
	t := itoa(target)

	var nums []string
	for i := 0; i < max; i++ {
		nums = append(nums, nextLineStrings()[0])
	}

	if sliceContains(nums, t) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func sliceContains[T comparable](slice []T, key T) bool {
	for _, s := range slice {
		if s == key {
			return true
		}
	}
	return false
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextLineInts() []int {
	line := nextLine()
	parts := strings.Fields(line)
	var nums []int
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func nextLineStrings() []string {
	line := nextLine()
	return strings.Fields(line)
}

// Q
// 長さ N の重複した要素の無い数列 A と整数 K が与えられるので、
// A に K が含まれていれば "YES" を、そうでなければ "NO" を出力してください。

// ▼　下記解答欄にコードを記入してみよう

// 入力される値
// N K
// A_1
// ...
// A_N

// ・1 行目では、配列 A の要素数 N と検索する値 K が半角スペース区切りで与えられます。
// ・続く N 行では、配列 A の要素が先頭から順に与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// ・A に K が含まれていれば "YES" を、そうでなければ "NO" を 1 行で出力してください。
// ・また、出力の末尾には改行を入れてください。

// 条件
// ・1 ≦ N ≦ 100,000
// ・0 ≦ A_i , K ≦ 1,000,000 (1 ≦ i ≦ N)

// 入力例1
// 3 5
// 1
// 3
// 5

// 出力例1
// YES

// 入力例2
// 5 4
// 1
// 2
// 3
// 5
// 6

// 出力例2
// NO
