package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Hero struct {
	ID     int // ID
	Level  int // レベル
	HP     int // 体力
	Power  int // 攻撃力
	Guard  int // 防御力
	Speed  int // 素早さ
	Wisdom int // 賢さ
	Luck   int // 運
}

type EventType string

const (
	EventTypeLevelUp        EventType = "levelup"
	EventTypeMuscleTraining EventType = "muscle_training"
	EventTypeRunning        EventType = "running"
	EventTypeStudy          EventType = "study"
	EventTypePray           EventType = "pray"
)

func NewHero(id, level, hp, power, guard, speed, wisdom, luck int) *Hero {
	return &Hero{
		ID:     id,
		Level:  level,
		HP:     hp,
		Power:  power,
		Guard:  guard,
		Speed:  speed,
		Wisdom: wisdom,
		Luck:   luck,
	}
}

func main() {
	infoArray := nextLineInts()
	heroCount := infoArray[0]
	events := infoArray[1]

	heroes := make(map[int]*Hero, heroCount)
	for i := 1; i <= heroCount; i++ {
		heroInfoArray := nextLineInts()
		heroes[i] = NewHero(i, heroInfoArray[0], heroInfoArray[1], heroInfoArray[2], heroInfoArray[3], heroInfoArray[4], heroInfoArray[5], heroInfoArray[6])
	}

	for i := 0; i < events; i++ {
		eventInfoArray := nextLineStrings()
		heroID := atoi(eventInfoArray[0])
		hero, ok := heroes[heroID]
		if !ok {
			panic("Hero not found")
		}
		event := EventType(eventInfoArray[1])
		switch event {
		case EventTypeLevelUp:
			hero.Level += 1
			hero.HP += atoi(eventInfoArray[2])
			hero.Power += atoi(eventInfoArray[3])
			hero.Guard += atoi(eventInfoArray[4])
			hero.Speed += atoi(eventInfoArray[5])
			hero.Wisdom += atoi(eventInfoArray[6])
			hero.Luck += atoi(eventInfoArray[7])
		case EventTypeMuscleTraining:
			hero.HP += atoi(eventInfoArray[2])
			hero.Power += atoi(eventInfoArray[3])
		case EventTypeRunning:
			hero.Guard += atoi(eventInfoArray[2])
			hero.Speed += atoi(eventInfoArray[3])
		case EventTypeStudy:
			hero.Wisdom += atoi(eventInfoArray[2])
		case EventTypePray:
			hero.Luck += atoi(eventInfoArray[2])
		}
	}

	for i := 1; i <= heroCount; i++ {
		hero := heroes[i]
		fmt.Println(hero.Level, hero.HP, hero.Power, hero.Guard, hero.Speed, hero.Wisdom, hero.Luck)
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
// paiza 村にたびたび魔物が訪れるため、 1 〜 N 番の番号が割り当てられた N 人の勇者を雇うことにしました。
// 勇者には次のようなステータスを持ちます。

// レベル l_i
// 体力 h_i
// 攻撃力 a_i
// 防御力 d_i
// 素早さ s_i
// 賢さ c_i
// 運 f_i

// また、各勇者には次のようなイベントが発生します。

// levelup h a d s c f
// レベルが 1 上昇
// 体力が h 上昇
// 攻撃力が a 上昇
// 防御力が d 上昇
// 素早さが s 上昇
// 賢さが c 上昇
// 運が f 上昇

// muscle_training h a
// 体力が h 上昇
// 攻撃力が a 上昇

// running d s
// 防御力が d 上昇
// 素早さが s 上昇

// study c
// 賢さが c 上昇

// pray f
// 運が f 上昇

// 各勇者の初期ステータスと起こるイベントの回数、
// また、起こるイベントとその対象の勇者の番号が与えられるので、
// 全てのイベントが終わった後の各勇者のステータスを出力してください。

// Input1
// 1 3
// 23 128 533 552 44 69 420
// 1 muscle_training 565 241
// 1 study 132
// 1 levelup 379 585 4 145 276 8

// Output1
// 24 1072 1359 556 189 477 428

// Input2
// 10 20
// 161 295 842 678 857 640 702
// 703 973 816 584 474 996 452
// 640 929 296 484 617 785 968
// 621 946 565 298 297 17 963
// 82 75 684 44 706 828 615
// 509 27 178 957 156 705 150
// 224 247 745 338 11 969 218
// 343 25 757 600 277 478 814
// 217 537 596 50 807 363 299
// 123 296 770 108 25 500 938
// 4 muscle_training 367 195
// 8 pray 229
// 10 levelup 683 829 497 446 843 38
// 3 pray 505
// 6 pray 488
// 6 muscle_training 280 653
// 4 study 630
// 10 muscle_training 111 609
// 8 levelup 846 819 872 906 126 58
// 7 muscle_training 75 112
// 3 levelup 750 656 95 557 50 95
// 7 study 771
// 3 muscle_training 251 458
// 8 study 888
// 4 study 52
// 3 pray 912
// 10 study 264
// 2 pray 886
// 5 muscle_training 1000 676
// 9 study 125

// Output2
// 161 295 842 678 857 640 702
// 703 973 816 584 474 996 1338
// 641 1930 1410 579 1174 835 2480
// 621 1313 760 298 297 699 963
// 82 1075 1360 44 706 828 615
// 509 307 831 957 156 705 638
// 224 322 857 338 11 1740 218
// 344 871 1576 1472 1183 1492 1101
// 217 537 596 50 807 488 299
// 124 1090 2208 605 471 1607 976
