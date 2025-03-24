package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	targetNumber := 100

	var primeNumbers []int
	for i := 2; i <= targetNumber; i++ {
		if isPrimeNumber(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}

	fmt.Println(primeNumbers)
}

func isPrimeNumber(target int) bool {
	if target < 2 {
		return false
	}

	// √iまでで割り切れるかどうかを確認するだけで良い
	for i := 2; i*i < target; i++ {
		if target%i == 0 {
			return false
		}
	}

	return true
}

// Output
// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
