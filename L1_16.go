package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := []int{}
	right := []int{}

	for _, num := range arr[:len(arr)-1] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func main() {
	arr := []int{1, 2, -1, 0, 17, -42, 4, 1, 4}

	fmt.Println("before:", arr)
	arr = quicksort(arr)
	fmt.Println("after:", arr)

}
