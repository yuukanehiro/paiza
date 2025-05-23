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
	targetNum := infoArray[1]
	insertNum := infoArray[2]

	nums := []int{}
	for i := 0; i < max; i++ {
		infoArray := nextLineInts()
		nums = append(nums, infoArray[0])
	}

	// firstHalf := nums[:targetNum]
	// secondHalf := nums[targetNum:]
	// secondHalf = append([]int{insertNum}, secondHalf...)
	// firstHalf = append(firstHalf, secondHalf...)

	// 上記と同じ良い書き方
	nums = append(nums[:targetNum], append([]int{insertNum}, nums[targetNum:]...)...)

	for _, num := range nums {
		fmt.Println(num)
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
// 整数 N, K, Q と、 長さ N の配列 A_1, A_2, ..., A_N が与えられるので、
// A_K の後ろに Q を挿入した後の長さ N+1 の配列について、先頭から改行区切りで出力してください。

// 入力される値
// N K Q
// A_1
// ...
// A_N

// ・1 行目では、配列 A の要素数 N と整数 K , Q が半角スペース区切りで与えられます。
// ・続く N 行では、配列 A の要素が先頭から順に与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// A_1
// ...
// A_{N+1}

// ・Q を A_K の後ろに挿入した後の配列の各要素を先頭から改行区切りで出力してください。
// ・また、出力の末尾には改行を入れてください。
// 条件
// ・1 ≦ N ≦ 100,000
// ・1 ≦ K ≦ N
// ・0 ≦ Q ≦ 100
// ・0 ≦ A_i ≦ 100 (1 ≦ i ≦ N)

// 入力例1
// 3 1 57
// 17
// 57
// 83

// 出力例1
// 17
// 57
// 57
// 83

// 入力例2
// 10 6 45
// 38
// 83
// 46
// 57
// 15
// 30
// 51
// 88
// 96
// 85

// 出力例2
// 38
// 83
// 46
// 57
// 15
// 30
// 45
// 51
// 88
// 96
// 85
