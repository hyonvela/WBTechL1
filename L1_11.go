package main

import (
	"fmt"
)

func main() {
	// Множество реализуется в виде мапы, в которой ключи - элементы множества
	// Значения struct{} выбраны потому что они занимают 0 байт
	setA := map[int]struct{}{
		1: {},
		0: {},
		3: {},
		2: {},
	}

	setB := map[int]struct{}{
		3: {},
		5: {},
		4: {},
		2: {},
	}

	intersection := make(map[int]struct{})

	for key := range setA {
		if _, exists := setB[key]; exists {
			intersection[key] = struct{}{}
		}
	}

	fmt.Println("Пересечение множеств:", intersection)
}
