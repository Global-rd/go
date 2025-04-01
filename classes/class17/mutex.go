package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

func MutexExample() {
	counter := Counter{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			counter.Lock()

			counter.value++

			counter.Unlock()
		}(i)
	}
	time.Sleep(time.Second)

	counter.Lock()

	fmt.Println("counter", counter.value)

	counter.Unlock()
}
