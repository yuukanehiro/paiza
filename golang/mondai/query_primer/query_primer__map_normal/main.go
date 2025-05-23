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

	users := make(map[int]string)
	for i := 0; i < max; i++ {
		userInfoArray := nextLineStrings()
		id := atoi(userInfoArray[0])
		users[id] = userInfoArray[1]
	}

	var cmds [][]string
	for i := 0; i < cmdCount; i++ {
		cmd := nextLineStrings()
		cmds = append(cmds, cmd)
	}

	for _, cmd := range cmds {
		if cmd[0] == "join" {
			id := atoi(cmd[1])
			name := cmd[2]
			users[id] = name
		} else if cmd[0] == "leave" {
			id := atoi(cmd[1])
			delete(users, id)
		} else if cmd[0] == "call" {
			id := atoi(cmd[1])
			if name, ok := users[id]; ok {
				fmt.Println(name)
			} else {
				panic("not found")
			}
		} else if cmd[0] == "join" {
			id := atoi(cmd[1])
			name := cmd[2]
			users[id] = name
		} else {
			panic("Unknown command")
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
// 3xxx年、ロボット学校の先生である paiza 君は、新しく担当するクラスの生徒一人一人の出席番号と識別 ID を覚えて、出席番号が与えられたら、その生徒の識別 ID を言えるようになる必要があります。
// paiza 君の務める学校は転校が多く、頻繁に生徒が増減します。

// 覚えるべき生徒の出席番号と識別 ID が与えられたのち、いくつかのイベントを表す文字列が与えられるので、与えられた順に各イベントに応じて次のような処理をおこなってください。

// ・join num id
// 生徒番号 num , 識別ID id の生徒を新たに覚える

// ・leave num
// 生徒番号 num の生徒を忘れる

// ・call num
// 生徒番号 num の生徒の識別 ID を出力する

// 入力される値
// N K
// num_1 ID_1
// ...
// num_N ID_N
// S_1
// ...
// S_K

// ・1 行目では、初めに覚える生徒の人数 N と与えられるイベントの回数 K が与えられます。
// ・続く N 行のうち i 行目 (1 ≦ i ≦ N) では、i 番目の生徒の出席番号と識別 ID の組 num_i , ID_i が半角スペース区切りで与えられます。
// ・続く K 行では、起きるイベントを表す文字列 S_i (1 ≦ i ≦ K) が与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// ・イベントに従って出力してください。
// ・また、出力の末尾には改行を入れてください。

// 条件
// ・1 ≦ N , K ≦ 100,000
// ・1 ≦ num_i ≦ 1,000,000 (1 ≦ i ≦ N)
// ・num_i ≠ num_j (i ≠ j)
// ・ID_i は アルファベット大文字小文字と数字から成る 20 文字以下の文字列 (1 ≦ i ≦ N)
// ・S_i は次のいずれかの形式である。

// ・join num id
// 生徒番号 num , 識別 ID id の生徒を新たに覚える。

// ・leave num
// 生徒番号 num の生徒を忘れる。

// ・call num
// 生徒番号 num の生徒の識別 ID を出力する。
// この時点で生徒番号 num の生徒がいることは保証されている。

// 1 ≦ num ≦ 1,000,000
// id は 20 文字以下の文字列
// 入力例1
// 4 4
// 1 Sin
// 2 Sakura
// 3 Kayo
// 4 Yui
// call 4
// leave 2
// join 2 Sakuya
// call 2

// 出力例1
// Yui
// Sakuya

// 入力例2
// 5 10
// 696042 pieF4
// 162082 Geig1
// 43482 Ich7D
// 647458 foh8C
// 71317 Aiv4g
// call 43482
// call 696042
// call 696042
// leave 696042
// call 647458
// call 647458
// call 162082
// join 591845 Ue7wo
// call 591845
// leave 647458

// 出力例2
// Ich7D
// pieF4
// pieF4
// foh8C
// foh8C
// Geig1
// Ue7wo
