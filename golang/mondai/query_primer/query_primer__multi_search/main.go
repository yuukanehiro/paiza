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
	targetCount := infoArray[1]

	nums := map[int]bool{}
	for i := 0; i < max; i++ {
		n := nextLineInts()[0]
		nums[n] = true
	}
	var targets []int
	for i := 0; i < targetCount; i++ {
		targets = append(targets, nextLineInts()[0])
	}

	for _, target := range targets {
		if mapContains(nums, target) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func mapContains[T comparable](set map[T]bool, key T) bool {
	_, exists := set[key]
	return exists
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
// 長さ N の重複した要素の無い数列 A と Q 個の整数 K_1 ... K_Q が与えられるので、
// 各 K_i について、 A に K_i が含まれていれば "YES" を、そうでなければ "NO" を出力してください。

// ▼　下記解答欄にコードを記入してみよう

// 入力される値
// N Q
// A_1
// ...
// A_N
// K_1
// ...
// K_Q

// ・1 行目では、配列 A の要素数 N と検索する値の個数 Q が半角スペース区切りで与えられます。
// ・続く N 行では、配列 A の要素が先頭から順に与えられます。
// ・続く Q 行では、検索する値 K_1 .. K_Q が順に与えられます

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// ・Q 行出力してください。i 行目には A に K_i が含まれていれば "YES" を、そうでなければ "NO" を出力してください。
// ・また、出力の末尾には改行を入れてください。

// 条件
// ・1 ≦ N , Q ≦ 100,000
// ・0 ≦ A_i ≦ 1,000,000 (1 ≦ i ≦ N)
// ・0 ≦ K_i ≦ 1,000,000 (1 ≦ i ≦ Q)

// 入力例1
// 5 5
// 1
// 2
// 3
// 4
// 5
// 1
// 3
// 5
// 7
// 9

// 出力例1
// YES
// YES
// YES
// NO
// NO

// 入力例2
// 10 5
// 351051
// 62992
// 166282
// 497610
// 636807
// 678131
// 885162
// 81763
// 810110
// 943644
// 670661
// 463229
// 62992
// 1973
// 901393

// 出力例2
// NO
// NO
// YES
// NO
// NO
