package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	var mu sync.Mutex

	sum := 0

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			mu.Lock()
			sum += num * num
			mu.Unlock()

			defer wg.Done()
		}(num)
	}

	wg.Wait()

	fmt.Println(sum)
}
