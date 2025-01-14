package main

import "fmt"

func SetBit(num int64, i int, val int) int64 {
	if val == 1 {
		// Устанавливаем i-й бит в 1
		num |= (1 << i)
	} else {
		// Устанавливаем i-й бит в 0
		num &= ^(1 << i)
	}
	return num
}

func main() {
	var num int64 = 0
	var i int
	var val int

	fmt.Println("num = 0")
	fmt.Println("Enter bit index (i):")
	fmt.Scan(&i)

	if i < 0 || i > 64 {
		fmt.Println("Error. index must be in range [0, 64].")
		return
	}

	fmt.Println("Enter value (0 или 1):")
	fmt.Scan(&val)

	if val != 0 && val != 1 {
		fmt.Println("Error. value must be 0 or 1 only.")
		return
	}

	num = SetBit(num, i, val)

	fmt.Printf("New num: %d\n", num)
}
