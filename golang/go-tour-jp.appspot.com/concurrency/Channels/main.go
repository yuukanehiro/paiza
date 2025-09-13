package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // [7, 2, 8]を処理するgoroutine ... 17
	go sum(s[len(s)/2:], c) // [-9, 4, 0]を処理するgoroutine ... -5
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}

// Output:
// -5 17 12
