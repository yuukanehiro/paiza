package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	infoArray := nextLineInts()
	max := infoArray[0]
	heightX := infoArray[1]
	heightMe := infoArray[2]

	var heightArray []int
	for i := 0; i < max; i++ {
		heightArray = append(heightArray, nextLineInts()[0])
	}

	heightArray = append(heightArray, heightX)
	heightArray = append(heightArray, heightMe)
	sortAsc(heightArray)
	index := sort.SearchInts(heightArray, heightMe)

	fmt.Println(index + 1)
}

// 昇順ソート
func sortAsc(numbers []int) {
	sort.Ints(numbers)
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
