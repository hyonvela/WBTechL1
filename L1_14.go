package main

import "fmt"

func detectType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("type int, value", v)
	case string:
		fmt.Println("type string, value", v)
	case bool:
		fmt.Println("type bool, value", v)
	case chan int:
		fmt.Println("type channel, value", v)
	default:
		fmt.Println("unknown type, value", v)
	}
}

func main() {
	vars := []interface{}{
		10,
		"string",
		true,
		make(chan int),
	}

	for _, v := range vars {
		detectType(v)
	}
}
