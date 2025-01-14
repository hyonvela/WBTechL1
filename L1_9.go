package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	inputChannel := make(chan int)
	outputChannel := make(chan int)

	go func() {
		for _, num := range numbers {
			inputChannel <- num
			time.Sleep(1 * time.Second)
		}
		close(inputChannel)
	}()

	go func() {
		for num := range inputChannel {
			outputChannel <- num * 2
		}
		close(outputChannel)
	}()

	for result := range outputChannel {
		fmt.Println(result)
	}
}
