package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Println("a =", a, ", b =", b)

	fmt.Println("Arifmetic swapping")
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a =", a, ", b =", b)

	fmt.Println("XOR swapping")
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("a =", a, ", b =", b)

}
