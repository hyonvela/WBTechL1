package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	reversed := make([]rune, n)

	for i, r := range runes {
		reversed[n-1-i] = r
	}

	return string(reversed)
}

func main() {
	str := "главрыба"
	reversed := reverseString(str)
	fmt.Println(str)
	fmt.Println(reversed)
}
