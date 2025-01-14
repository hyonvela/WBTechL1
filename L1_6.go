package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker with context stopped")
			return
		default:
			fmt.Println("Working with context...")
			time.Sleep(1 * time.Second)
		}
	}
}

func workerWithChannel(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Worker with channel stopped")
			return
		default:
			fmt.Println("Working with channel...")
			time.Sleep(1 * time.Second)
		}
	}
}

var stop bool
var mu sync.Mutex

func workerWithFlag() {
	for {
		mu.Lock()
		if stop {
			mu.Unlock()
			fmt.Println("Worker with flag stopped")
			return
		}
		mu.Unlock()

		fmt.Println("Working with flag...")
		time.Sleep(1 * time.Second)
	}
}

func workerWithWaitGroup(wg *sync.WaitGroup, t int) {
	for i := 0; i < t; i++ {
		fmt.Println("Working with WaitGroup...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Worker with WaitGroup stopped")
	defer wg.Done()
}

func workerWithDefer() {
	for {
		fmt.Println("Working with defer...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Starting worker with context...")
	ctx, cancel := context.WithCancel(context.Background())
	go workerWithContext(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	fmt.Println("Starting worker with channel...")
	done := make(chan bool)
	go workerWithChannel(done)

	time.Sleep(5 * time.Second)
	done <- true
	time.Sleep(1 * time.Second)

	fmt.Println("Starting worker with flag...")
	go workerWithFlag()

	time.Sleep(5 * time.Second)
	mu.Lock()
	stop = true
	mu.Unlock()
	time.Sleep(1 * time.Second)

	fmt.Println("Starting worker with WaitGroup...")
	var wg sync.WaitGroup
	wg.Add(1)
	go workerWithWaitGroup(&wg, 5)
	wg.Wait()

	fmt.Println("Starting worker with defer...")
	go workerWithDefer()
	defer fmt.Println("Worker with defer stopped")
	time.Sleep(5 * time.Second)
}
