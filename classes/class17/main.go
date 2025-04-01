package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

func main() {
	var counter int64
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, int64(1))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("After addition:", counter)
}

func UintPtrExample() {
	var pointer unsafe.Pointer
	data := float64(1)
	atomic.StorePointer(&pointer, unsafe.Pointer(&data))

	loaded := atomic.LoadPointer(&pointer)
	fmt.Println("Loaded value:", *(*float64)(loaded))
}
