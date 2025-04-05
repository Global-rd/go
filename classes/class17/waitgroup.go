package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("working")
			time.Sleep(time.Second * 2)
		}()
	}

	fmt.Println("wait")

	wg.Wait()

	fmt.Println("finish")
}
