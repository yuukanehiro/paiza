package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Robot struct {
	ID    int
	Level int
	X     int // X座標
	Y     int // Y座標
}

func NewRobot(id int, x int, y int, level int) *Robot {
	return &Robot{
		ID:    id,
		Level: level,
		X:     x,
		Y:     y,
	}
}

func (r *Robot) Move(direction string) {
	var distance int
	switch r.Level {
	case 1:
		distance = 1
	case 2:
		distance = 2
	case 3:
		distance = 5
	case 4:
		distance = 10
	}

	switch direction {
	case "N":
		r.Y -= distance // 数学の座標とは逆
	case "S":
		r.Y += distance // 数学の座標とは逆
	case "E":
		r.X += distance
	case "W":
		r.X -= distance
	}
}

func (r *Robot) LevelUp() {
	if r.Level < 4 {
		r.Level++
	}
}

func main() {
	const levelUpPointCount = 10 // レベルアップするマスの数

	infoArray := nextLineInts()
	// yMax := infoArray[0]          // ロボットの初期位置のY座標の最大値 ... 利用しない
	// xMax := infoArray[1]          // ロボットの初期位置のX座標の最大値 ... 利用しない
	robotCount := infoArray[2]    // ロボットの数
	activityCount := infoArray[3] // アクティビティの数

	levelUpMap := make(map[int]map[int]bool)
	for i := 0; i < levelUpPointCount; i++ {
		levelUpPoint := nextLineInts()
		if levelUpMap[levelUpPoint[0]] == nil {
			levelUpMap[levelUpPoint[0]] = make(map[int]bool)
		}
		levelUpMap[levelUpPoint[0]][levelUpPoint[1]] = true
	}

	robots := make(map[int]*Robot, robotCount)
	for i := 1; i <= robotCount; i++ {
		robotInfoArray := nextLineInts()
		robots[i] = NewRobot(i, robotInfoArray[0], robotInfoArray[1], robotInfoArray[2])
	}

	for i := 0; i < activityCount; i++ {
		activityInfoArray := nextLineStrings()
		robotID := atoi(activityInfoArray[0])
		direction := activityInfoArray[1]

		if robots[robotID] == nil {
			panic("Robot not found")
		}

		robots[robotID].Move(direction)

		if levelUpMap[robots[robotID].X][robots[robotID].Y] {
			robots[robotID].LevelUp()
		}
	}

	for i := 1; i <= robotCount; i++ {
		robot := robots[i]
		fmt.Println(robot.X, robot.Y, robot.Level)
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
// paiza 株式会社では、物品の管理のために上の図のような座標系の広さが無限大のマスの工場 で 番号 1 〜 N が割り当てられた N 台のロボットを運用していました。ところがある日、全てのロボットが暴走してしまいました。各ロボットは性能ごとにレベル分けされており、次の通り移動距離が決まっています。

// Lv1 : 特定の方角に 1 マス進む
// Lv2 : 特定の方角に 2 マス進む
// Lv3 : 特定の方角に 5 マス進む
// Lv4 : 特定の方角に 10 マス進む

// また、工場のマスのうち 10 マスには工具箱が置かれており、移動後にそのマスにロボットがぴったり止まっていた場合、そのロボットのレベルが 1 上がってしまいます（最大レベルは 4)。
// レベル l のロボットの初期位置が工具箱の置かれているマスであったとしても、そのロボットのレベルは l で始まることに気をつけてください。

// 幸い、初めにロボットがいる範囲や工具箱の置かれているマス、各ロボットの位置とレベル、また、どのロボットがどのような順番でどの方角に移動するかの情報はわかっているので、ロボットの移動が K 回終わったときの各ロボットの位置とレベルを推定してください。

// ▼　下記解答欄にコードを記入してみよう

// 入力される値
// H W N K
// lx_1 ly_1
// ...
// lx_10 ly_10
// x_1 y_1 l_1
// ...
// x_N y_N l_N
// r_1 d_1
// ...
// r_K d_K

// ・ 1 行目では ロボットの初期位置の y , x 座標の上限についての整数 H , W , ロボットの数 N , ロボットの移動回数 K が半角スペース区切りで与えられます。
// ・ 続く 10 行のうち i 行目では、i 個目の工具箱が置かれたマスの x , y 座標 x_i , y_i が与えられます。(1 ≦ i ≦ 10)
// ・ 続く N 行のうち i 行目では、 番号 i のロボットの初期位置の x 座標 x_i , y 座標 y_i , レベル l_i が半角スペース区切りで与えられます。
// ・ 続く K 行のうち i 行目では、 i 回目の移動を行ったロボットの番号 r_i と移動の方角 d_i が与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// i 番のロボットの最終的な位置 x_i , y_i とレベル l_i を i 行目に出力してください。

// x_1 y_1 l_1
// ...
// x_N y_N l_N
// 条件
// ・ 5 ≦ H , W , N , K ≦ 10^5
// ・ 0 ≦ lx_i < W , 0 ≦ ly_i < H (1 ≦ i ≦ 10)
// ・ 0 ≦ x_i < W , 0 ≦ y_i < H , 1 ≦ l_i ≦ 4 (1 ≦ i ≦ N)
// ・ 1 ≦ r_i ≦ N
// ・ d_i は "N" , "S" , "E" , "W" のいずれか (1 ≦ i ≦ K) で、それぞれ 北・南・東・西 へ移動したことを表す。

// 入力例1
// 5 5 3 3
// 0 0
// 0 1
// 0 2
// 0 3
// 0 4
// 1 0
// 1 1
// 1 2
// 1 3
// 1 4
// 2 1 1
// 2 2 1
// 2 3 1
// 1 W
// 1 E
// 3 S

// 出力例1
// 3 1 2
// 2 2 1
// 2 4 1
