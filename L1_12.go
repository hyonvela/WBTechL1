package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	// Множество реализуется в виде мапы, в которой ключи - элементы множества
	// Значения struct{} выбраны потому что они занимают 0 байт
	set := make(map[string]struct{})

	for _, str := range strs {
		set[str] = struct{}{}
	}

	fmt.Println(set)
}
