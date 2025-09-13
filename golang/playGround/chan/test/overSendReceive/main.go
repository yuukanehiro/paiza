package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // ここで受信しようとしているが、送信されていないため、デッドロックになる
}

// Output:
// % go run main.go
// 1
// 2
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /Users/kanehiroyuu/Documents/GitHub/paiza/golang/playGround/chan/test/overSendReceive/main.go:12 +0xe8
// exit status 2
