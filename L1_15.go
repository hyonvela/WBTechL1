package main

import (
	"fmt"
	"math/rand"
)

func createHugeString(size int) string {
	chars := []rune(" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")
	result := make([]rune, size)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	// Создаем новую строку, чтобы избежать утечки памяти
	newStr := string([]byte(v[:100]))
	return newStr
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
