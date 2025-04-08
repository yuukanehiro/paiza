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
	hubCount := infoArray[0]
	activityCount := infoArray[1]
	startHub := infoArray[2]

	hubMap := make(map[int]map[int]string)
	for i := 1; i <= hubCount; i++ {
		hubMap[i] = make(map[int]string)
		hubInfoArray := nextLineStrings()
		hubMap[i][0] = hubInfoArray[0] // 文字
		hubMap[i][1] = hubInfoArray[1] // 道1
		hubMap[i][2] = hubInfoArray[2] // 道2
	}
	xActivities := []int{}
	for i := 0; i < activityCount; i++ {
		activityInfoArray := nextLineInts()
		xActivities = append(xActivities, activityInfoArray[0])
	}

	answerStrings := []string{}
	nextHub := "0"
	for i := 0; i < len(xActivities); i++ {
		if i == 0 {
			answerStrings = append(answerStrings, hubMap[startHub][0])
			nextHub = hubMap[startHub][xActivities[0]]
			continue
		}
		answerStrings = append(answerStrings, hubMap[atoi(nextHub)][0])
		nextHub = hubMap[atoi(nextHub)][xActivities[i]]

		if i == len(xActivities)-1 {
			answerStrings = append(answerStrings, hubMap[atoi(nextHub)][0])
			break
		}
	}

	res := strings.Join(answerStrings, "") // 区切り文字なしで連結
	fmt.Println(res)

	// b, _ := json.MarshalIndent(hubMap, "", "  ")
	// println(string(b))
	// {
	// 	"1": {
	// 	  "0": "p",
	// 	  "1": "2",
	// 	  "2": "4"
	// 	},
	// 	"2": {
	// 	  "0": "a",
	// 	  "1": "3",
	// 	  "2": "1"
	// 	},
	// 	"3": {
	// 	  "0": "i",
	// 	  "1": "4",
	// 	  "2": "2"
	// 	},
	// 	"4": {
	// 	  "0": "z",
	// 	  "1": "1",
	// 	  "2": "2"
	// 	}
	//   }
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
// 洞窟を探検していたあなたは出口のない迷路に迷い込んでしまいました。
// 脱出するには、迷路の地点を与えられた指示通りに移動し、移動で訪れた（移動の開始・終了地点を含む）地点に置かれたアルファベットを
// つなげた文字列を呪文として唱える必要があります。
// 各頂点からは、他の頂点に向かって一方通行の 2 つの道が伸びています。
// 各ポイントの情報と移動の指示が与えられるので、呪文となる文字列を出力してください。

// Input1
// 4 4 1
// p 2 4
// a 3 1
// i 4 2
// z 1 2
// 1
// 1
// 1
// 2

// Output1
// paiza

// Input2
// 5 10 5
// o 5 4
// f 1 5
// b 1 2
// k 1 5
// k 2 4
// 1
// 2
// 1
// 2
// 2
// 2
// 2
// 2
// 1
// 1

// Output2
// kfkfkkkkkfo
