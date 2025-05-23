package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type CarI interface {
	Run()
	Fly()
	Teleport()
	GetDistance() int
}

type BaseCar struct {
	Fuel           int
	FuelEfficiency int
	Distance       int
}

func (b *BaseCar) GetDistance() int {
	return b.Distance
}

type SuperCar struct {
	BaseCar
}

func (s *SuperCar) Run() {
	if s.Fuel < 1 {
		return
	}
	s.Fuel -= 1
	s.Distance += s.FuelEfficiency
}

func (s *SuperCar) Fly()      {}
func (s *SuperCar) Teleport() {}

type SuperSuperCar struct {
	BaseCar
}

func (s *SuperSuperCar) Run() {
	if s.Fuel < 1 {
		return
	}
	s.Fuel -= 1
	s.Distance += s.FuelEfficiency
}
func (s *SuperSuperCar) Fly() {
	if s.Fuel < 5 {
		s.Run()
		return
	}
	s.Fuel -= 5
	s.Distance += s.FuelEfficiency * s.FuelEfficiency
}
func (s *SuperSuperCar) Teleport() {}

type SuperSuperSuperCar struct {
	BaseCar
}

func (s *SuperSuperSuperCar) Run() {
	if s.Fuel < 1 {
		return
	}
	s.Fuel -= 1
	s.Distance += s.FuelEfficiency
}
func (s *SuperSuperSuperCar) Fly() {
	if s.Fuel < 5 {
		s.Run()
		return
	}
	s.Fuel -= 5
	s.Distance += 2 * s.FuelEfficiency * s.FuelEfficiency
}
func (s *SuperSuperSuperCar) Teleport() {
	if s.Fuel < s.FuelEfficiency*s.FuelEfficiency {
		s.Fly()
		return
	}
	s.Fuel -= s.FuelEfficiency * s.FuelEfficiency
	s.Distance += s.FuelEfficiency * s.FuelEfficiency * s.FuelEfficiency * s.FuelEfficiency
}

func main() {
	infoArray := nextLineInts()
	carCount := infoArray[0]
	activityCount := infoArray[1]

	cars := make(map[int]CarI, carCount)
	for i := 1; i <= carCount; i++ {
		carInfoArray := nextLineStrings()
		switch carInfoArray[0] {
		case "supercar":
			cars[i] = &SuperCar{
				BaseCar: BaseCar{
					Fuel:           atoi(carInfoArray[1]),
					FuelEfficiency: atoi(carInfoArray[2]),
				},
			}
		case "supersupercar":
			cars[i] = &SuperSuperCar{
				BaseCar: BaseCar{
					Fuel:           atoi(carInfoArray[1]),
					FuelEfficiency: atoi(carInfoArray[2]),
				},
			}
		case "supersupersupercar":
			cars[i] = &SuperSuperSuperCar{
				BaseCar: BaseCar{
					Fuel:           atoi(carInfoArray[1]),
					FuelEfficiency: atoi(carInfoArray[2]),
				},
			}
		default:
			panic("Invalid car type")
		}
	}

	for i := 0; i < activityCount; i++ {
		activityInfoArray := nextLineStrings()
		carID := atoi(activityInfoArray[0])
		car := cars[carID]
		switch activityInfoArray[1] {
		case "run":
			car.Run()
		case "fly":
			car.Fly()
		case "teleport":
			car.Teleport()
		default:
			panic("Invalid activity type")
		}
	}

	for i := 1; i <= carCount; i++ {
		car := cars[i]
		fmt.Printf("%s\n", itoa(car.GetDistance()))
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
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
// よくクラスの題材を扱う際に、「クラスは車の設計書」といった例が出てきます。
// スーパーカー販売店に勤務しながらクラスの勉強をしていた paiza 君はスーパーカーの走る様子をクラスを用いてシミュレーションしてみようと考えました。
// ただ車を走らせてもつまらないので、陸を走るスーパーカーに加えて、空を飛べるスーパースーパーカー ・ テレポートできるスーパースーパースーパーカー もシミュレーションに加えた
// 番号 1 〜 N の N 台のシミュレーションをすることにしました。

// それぞれの車について、初めに入っている燃料の量 l と燃費 f が定まっており、加えて、車種に応じて次のような機能を持ちます。

// ・スーパーカー
// run
// 燃料を 1 消費し、 f (km) 走る。
// 燃料が 0 の場合は何も起こらない。

// ・スーパースーパーカー
// run
// 燃料を 1 消費し、 f (km) 走る。
// 燃料が 0 の場合は何も起こらない。

// fly
// 燃料を 5 消費し、 f^2 (km) 飛行する。
// 燃料が 5 未満の場合は run を行う。

// ・スーパースーパースーパーカー
// run
// 燃料を 1 消費し、 f (km) 走る。
// 燃料が 0 の場合は何も起こらない。

// fly
// 燃料を 5 消費し、 2 * f^2 (km) 飛行する。
// 燃料が 5 未満の場合は run を行う。

// teleport
// 燃料を f^2 消費し、 f^4 (km) 移動する。
// 燃料が f^2 未満の場合は fly を行う。

// シミュレートする車の台数 N と機能を使う回数 K , N 台の車の車種と機能を使った車の番号と使った機能が与えられるので、全てのシミュレーションが終わった後の、各車ごとの総移動距離を求めてください。

// ▼　下記解答欄にコードを記入してみよう

// 入力される値
// N K
// k_1 l_1 f_1
// ...
// k_N l_N f_N
// n_1 func_1
// ...
// n_K func_K

// ・ 1 行目では、シミュレートする車の台数 N と機能を使う回数 K が半角スペース区切りで与えられます。
// ・ 続く N 行のうち i 行目(1 ≦ i ≦ N)では、 i 番の車の種類 k_i , 初めに入っている燃料 l_i , 燃費 f_i が半角スペース区切りで与えられます。
// ・ 続く K 行では、車の番号 n_i と、使用するその車の機能 func_i が時系列順に与えられます。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// 全てのシミュレーションを終えたときの i 番の車の総移動距離 len_i を以下の形式で全ての車について出力してください。

// len_1
// ...
// len_N
// 条件
// ・ 与えられる値は全て整数
// ・ 1 ≦ N , K ≦ 10^5
// ・ k_i (1 ≦ i ≦ N) は "supercar","supersupercar","supersupersupercar" のいずれか
// ・ 1 ≦ l_i ≦ 10^5 , 1 ≦ f_i ≦ 100(1 ≦ i ≦ N)
// ・ 1 ≦ n_i ≦ N (1 ≦ i ≦ K)
// ・ func_i (1 ≦ i ≦ K) は "run" , "fly" , "teleport" のいずれか
// ・ 未定義の機能が呼び出されることはないことが保証されている

// 入力例1
// 3 6
// supercar 1 1
// supersupercar 10 10
// supersupersupercar 100 5
// 1 run
// 2 run
// 2 fly
// 3 run
// 3 fly
// 3 teleport

// 出力例1
// 1
// 110
// 680

// 入力例2
// 5 10
// supersupercar 1102 67
// supersupercar 63296 25
// supersupersupercar 47388 32
// supersupercar 30968 68
// supersupercar 53668 78
// 2 run
// 3 teleport
// 1 fly
// 2 run
// 4 run
// 5 fly
// 5 run
// 2 fly
// 4 run
// 1 fly

// 出力例2
// 8978
// 675
// 1048576
// 136
// 6162
