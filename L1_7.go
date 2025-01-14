package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu    sync.Mutex
	store map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		store: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("Key%d", i), i)
			time.Sleep(time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("All operations completed.")
	fmt.Println(safeMap.store)
}
