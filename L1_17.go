package main

import "fmt"

func binarySearch(arr []int, target int) int {
	l := arr[0]
	r := len(arr) - 1

	for l <= r {
		m := (l + r) / 2
		if arr[m] == target {
			return m
		}

		if target < arr[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 9

	fmt.Println("Array:", arr)
	fmt.Println("Target:", target)
	fmt.Println("Target's index:", binarySearch(arr, target))
}
