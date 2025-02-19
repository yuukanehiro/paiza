package main

func main() {
	var array = []int{10, 2, 3, 1, 7, 5, 6, 4, 9, 8}

	heapSort(array)

	for _, v := range array {
		println(v)
	}
}

func heapify(array []int, n, i int) {
	var largest = i
	var left = 2*i + 1
	var right = 2*i + 2

	if left < n && array[left] > array[largest] {
		largest = left
	}

	if right < n && array[right] > array[largest] {
		largest = right
	}

	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		// 元々のiの位置にあった値が、largestの位置に移動したので、再帰的にheapifyを呼び出す
		// これにより、最大値が最上位に来る
		// もし子ノードがない場合は、再帰的に呼び出されない
		heapify(array, n, largest)
	}
}

func heapSort(array []int) {
	var n = len(array)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}

	for i := n - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, i, 0)
	}
}

// % go run main.go
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
