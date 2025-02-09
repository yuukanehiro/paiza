package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Printf("array[0]: %d\n", array[0]) // 1
	fmt.Printf("array[1]: %d\n", array[1]) // 2

	// array[start:end] ... arrayのstart番目からend-1番目までの要素を取得
	fmt.Printf("array[1:3]: %v\n", array[1:3]) // [2 3] (array[1]からarray[3 - 1]まで)

	// (array[2]から最後まで)
	fmt.Printf("array[2:]: %v\n", array[2:]) // [3 4 5]

	// (array[0]からarray[2]まで)
	fmt.Printf("array[:3]: %v\n", array[:3]) // [1 2 3]

	fmt.Printf("array[1:len(array)]: %v\n", array[1:len(array)]) // [2 3 4 5]

	// (array[0]から最後の一つ手前まで)
	fmt.Printf("array[:len(array)-1]: %v\n", array[:len(array)-1]) // [1 2 3 4]

	// 全て
	fmt.Printf("array[:]: %v\n", array[:])                       // [1 2 3 4 5]
	fmt.Printf("array[0:len(array)]: %v\n", array[0:len(array)]) // [1 2 3 4 5]
}
