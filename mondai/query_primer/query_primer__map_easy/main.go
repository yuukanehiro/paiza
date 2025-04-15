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
	idCount := infoArray[1]

	users := make(map[int]string)
	for i := 0; i < max; i++ {
		userInfoArray := nextLineStrings()
		id := atoi(userInfoArray[0])
		users[id] = userInfoArray[1]
	}

	for i := 0; i < idCount; i++ {
		id := atoi(nextLine())
		if name, ok := users[id]; ok {
			fmt.Println(name)
		} else {
			panic("ID not found")
		}
	}
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}

	return i
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
// 3xxx 年、ロボット学校の先生である paiza 君は、新しく担当するクラスの生徒一人一人の出席番号と識別 ID を覚えるように言われました。
// 具体的には、出席番号が与えられたら、その生徒の識別 ID を言えるようになる必要があります。
// 覚えるべき生徒の出席番号と識別 ID のペアが与えられたのち、いくつか出席番号が与えられるので、各番号に対応する生徒の識別 ID を出力してください。

// 入力される値
// N K
// num_1 ID_1
// ...
// num_N ID_N
// Q_1
// ...
// Q_K

// ・1 行目では、生徒の人数 N と与えられる出席番号の個数 K が与えられます。
// ・続く N 行のうち i 行目 (1 ≦ i ≦ N) では、i 番目の生徒の出席番号と識別 ID の組 num_i , ID_i が半角スペース区切りで与えられます。
// ・続く K 行では、出席番号 Q_i (1 ≦ i ≦ K) が与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// A_1
// ...
// A_K

// ・出席番号が Q_i 番の生徒の識別 ID A_i を i 行目に出力してください。
// ・また、出力の末尾には改行を入れてください。
// 条件
// ・1 ≦ N , K ≦ 1,000
// ・0 ≦ num_i ≦ 10,000 (1 ≦ i ≦ N)
// ・num_i ≠ num_j (i ≠ j)
// ・ID_i は アルファベット大文字小文字(a ~ z , A ~ Z)と数字(0 ~ 9)から成る 20 文字以下の文字列 (1 ≦ i ≦ N)
// ・全ての Q_i について、その番号を出席番号として持つ生徒が必ず存在する。

// 入力例1
// 4 2
// 1 Sin
// 2 Sakura
// 3 Kayo
// 4 Yui
// 4
// 2

// 出力例1
// Yui
// Sakura

// 入力例2
// 10 5
// 5225 Eith5
// 5903 OoPi1
// 3824 ooZ0l
// 1979 cho4S
// 4319 Xa2ja
// 3371 Lo6em
// 5975 ceoZ0
// 7166 Ohz5A
// 8942 oi0Th
// 485 Qua2i
// 3371
// 3371
// 3824
// 3824
// 4319

// 出力例2
// Lo6em
// Lo6em
// ooZ0l
// ooZ0l
// Xa2ja
