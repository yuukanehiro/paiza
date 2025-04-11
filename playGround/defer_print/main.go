package main

import "fmt"

func main() {
	fmt.Println("start")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("end")
}

// Output:
// start
// end
// defer 3
// defer 2
// defer 1
