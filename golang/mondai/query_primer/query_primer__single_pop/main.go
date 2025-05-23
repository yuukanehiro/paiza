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
	n := infoArray[0]

	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, nextLineInts()[0])
	}

	res := nums[1:]

	for _, v := range res {
		fmt.Println(v)
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
// 数列 A が与えられるので、A の先頭の要素を削除した後の A を出力してください。

// ▼　下記解答欄にコードを記入してみよう

// 入力される値
// N
// A_1
// ...
// A_N

// ・1 行目では、配列 A の要素数 N が与えられます。
// ・続く N 行では、配列 A の要素が先頭から順に与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// A_2
// ...
// A_N

// ・A の先頭の要素を削除した後の A の各要素を先頭から順に改行区切りで出力してください。
// ・また、出力の末尾には改行を入れてください。
// 条件
// ・2 ≦ N ≦ 100,000
// ・0 ≦ A_i ≦ 10,000 (1 ≦ i ≦ N)

// 入力例1
// 10
// 5980
// 1569
// 5756
// 9335
// 9680
// 4571
// 5309
// 8696
// 9680
// 8963

// 出力例1
// 1569
// 5756
// 9335
// 9680
// 4571
// 5309
// 8696
// 9680
// 8963

// 入力例2
// 2
// 6963
// 9374

// 出力例2
// 9374
