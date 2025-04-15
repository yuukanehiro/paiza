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
	cmdCount := infoArray[1]

	var nums []int
	for i := 0; i < max; i++ {
		nums = append(nums, nextLineInts()[0])
	}

	var cmds []string
	for i := 0; i < cmdCount; i++ {
		cmds = append(cmds, nextLine())
	}

	for _, cmd := range cmds {
		if cmd == "pop" {
			if len(nums) > 0 {
				nums = nums[1:]
			}
		} else if cmd == "show" {
			for _, v := range nums {
				fmt.Println(v)
			}
		}
	}
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
// 数列 A と入力の回数 K が与えられるので、K 回の入力に応じて次のような処理をしてください。
// ・pop
// A の先頭の要素を削除する。既に A に要素が存在しない場合何もしない。
// ・show
// A の要素を先頭から順に改行区切りで出力する。A に要素が存在しない場合何も出力しない。

// 入力される値
// N K
// A_1
// ...
// A_N
// S_1
// ...
// S_K

// ・1 行目では、配列 A の要素数 N と与えられる入力の数 K が与えられます。
// ・続く N 行では、配列 A の要素が先頭から順に与えられます。
// ・続く K 行では、"pop" または "show" が与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// ・S_i で "show" が与えられる度に、A の全ての要素を先頭から順に改行区切りで出力してください。
// ・また、出力の末尾には改行を入れてください。

// 条件
// ・1 ≦ K ≦ N ≦ 100,000
// ・0 ≦ A_i ≦ 10,000 (1 ≦ i ≦ N)
// ・S_i (1 ≦ i ≦ K) は "pop" , "show" のいずれか
// ・S_i のうち、"show" であるものは 10 個以下であることが保証されている。

// 入力例1
// 5 3
// 7564
// 4860
// 2410
// 9178
// 7252
// pop
// pop
// show

// 出力例1
// 2410
// 9178
// 7252

// 入力例2
// 10 10
// 1005
// 2716
// 7856
// 8546
// 1339
// 4960
// 3926
// 9816
// 3018
// 4213
// pop
// pop
// pop
// pop
// show
// pop
// pop
// pop
// show
// pop

// 出力例2
// 1339
// 4960
// 3926
// 9816
// 3018
// 4213
// 9816
// 3018
// 4213
