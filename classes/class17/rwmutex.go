package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeData struct {
	value int
	mu    sync.RWMutex
}

func (s *SafeData) Read() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

func (s *SafeData) Write(val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value = val
}

func RWMutexExample() {
	data := SafeData{}
	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			data.Write(i)
			fmt.Println("Written:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				val := data.Read()
				fmt.Printf("Reader %d read: %d\n", id, val)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()

}
