package main


func sliceToMapSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool, len(slice))
	for _, v := range slice {
		set[v] = true
	}
	return set
}

func mapContains[T comparable](set map[T]bool, key T) bool {
	_, exists := set[key]
	return exists
}

// 使い方
// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	set := sliceToMapSet(slice)
// 	fmt.Println(mapContains(set, 3)) // true
