package main

// 絶対値を返却
// 例) 10 -> 10, -10 -> 10
func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// 素数を取得
func getPrimeNumbers(targetNumber int) []int {
	var primeNumbers []int

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}

		// √nで計算量を削減
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	for i := 0; i <= targetNumber; i++ {
		if isPrime(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}
