package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func worker(id int, ctx context.Context, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				// Канал закрыт, выходим из горутины
				return
			}
			fmt.Printf("Worker %d processed: %s\n", id, job)
		case <-ctx.Done():
			// контекст отменен, выходим из горутины
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Please provide a valid number of workers.")
		return
	}

	jobs := make(chan string)
	var wg sync.WaitGroup

	// использую context.WithCancel для создания контекста, который может быть отменен. Это позволяет корректно завершить все воркеры.
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ctx, jobs, &wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		fmt.Println("\nReceived interrupt signal. Shutting down...")
		cancel()    // Отмена контекста
		close(jobs) // Закрытие канала
	}()

	go func() {
		for i := 0; ; i++ {
			jobs <- fmt.Sprintf("Job %d", i)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("All workers have completed.")
}
