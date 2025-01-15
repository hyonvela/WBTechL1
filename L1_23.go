package main

import "fmt"

func removeElem(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Original:", slice)

	i := 3
	fmt.Println("Index:", i)

	slice = removeElem(slice, i)
	fmt.Println("New:", slice)
}
