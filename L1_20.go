package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	n := len(words)
	reversed := make([]string, n)

	for i, r := range words {
		reversed[n-1-i] = r
	}

	return strings.Join(reversed, " ")
}

func main() {
	str := "snow dog sun"
	reversed := reverseWords(str)
	fmt.Println(str)
	fmt.Println(reversed)
}
