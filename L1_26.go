package main

import (
	"fmt"
	"strings"
)

func hasUniqueCharacters(s string) bool {
	charMap := make(map[rune]bool)

	s = strings.ToLower(s)

	for _, char := range s {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

func main() {
	fmt.Println(hasUniqueCharacters("abcd"))
	fmt.Println(hasUniqueCharacters("abCdefAaf"))
	fmt.Println(hasUniqueCharacters("aabcd"))
}
