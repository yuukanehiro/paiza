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
	a := [][]string{
		{"1", "2", "3", "4"},
		{"10", "100", "0", "5"},
		{"8", "1", "3", "8"},
		{"15", "34", "94", "25"},
	}

	line := nextLine()
	var s []string
	s = strings.Split(line, " ")
	rowCount, _ := strconv.Atoi(s[0]) // 行数
	colCount, _ := strconv.Atoi(s[1]) // 列数

	fmt.Printf("%s", a[rowCount-1][colCount-1]) // 行数4, 列数3の場合は94
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
