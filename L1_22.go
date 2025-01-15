package main

import "fmt"

func main() {
	var a, b int64

	fmt.Println("Enter 2 numders > 2^20")
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("sum:", a+b)
	fmt.Println("dif:", a-b)
	fmt.Println("prod:", a*b)
	if b != 0 {
		fmt.Println("div:", a/b)
	} else {
		fmt.Println("Error. Cannot divide by 0")
	}
}
