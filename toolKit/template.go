package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
}

// 行を取得してstringで返却
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextLineInts() ([]int, error) {
	line := nextLine()
	parts := strings.Fields(line)
	var nums []int
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func nextLineStrings() ([]string, error) {
	line := nextLine()
	return strings.Fields(line), nil
}
