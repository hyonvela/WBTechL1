package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const duration = 10 * time.Second // Задайте нужное количество секунд

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		for i := 0; i < int(duration)/1000000000; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	go func() {
		for msg := range ch {
			fmt.Printf("Received: %d\n", msg)
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Println("Program finished after", duration)
}
