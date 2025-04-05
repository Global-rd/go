package main

import (
	"fmt"
	"time"
)

var DefaultTimeOut = time.Second * 2

func RunTimeOut() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		ch <- 1
	}()

	select {
	case v := <-ch:
		fmt.Println("value", v)
	case <-time.After(DefaultTimeOut):
		fmt.Println("timed out")
	}

	fmt.Println("finished")
}
