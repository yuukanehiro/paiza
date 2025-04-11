package main

import (
	"fmt"
)

func main() {
	var err error

	defer func() {
		fmt.Println("defer1:", err)
	}()

	defer func() {
		err = fmt.Errorf("error in defer3")
	}()

	defer func() {
		err = fmt.Errorf("error in defer2")
	}()

	err = fmt.Errorf("original error")
}

// Output:
// defer1: error in defer3
