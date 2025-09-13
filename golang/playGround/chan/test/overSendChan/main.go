package main

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // ここで送信しようとしているが、チャネルのバッファがいっぱいのため、デッドロックになる
}

// Output:

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//         /Users/kanehiroyuu/Documents/GitHub/paiza/golang/playGround/chan/test/overSendChan/main.go:7 +0x58
// exit status 2
