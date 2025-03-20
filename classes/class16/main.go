package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan User)
	// ch2 := make(chan int)

	go func() {
		for _, name := range []string{"Peter", "Tom", "Rob"} {
			ch1 <- User{
				Name: name,
			}
		}
		close(ch1)

		ch1 <- User{Name: "Anita"}
	}()

	ReadRange(ch1)

	// WriteChannels(ch1, ch2)
	// ReadChannels(ch1, ch2)
}

func ReadRange(ch <-chan User) {
	for u := range ch {
		fmt.Println("user", u)
	}
}

func WriteChannels(ch1 chan<- string, ch2 chan<- int) {
	go func() {
		ch1 <- "value1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 1
	}()
}

func ReadChannels(ch1 <-chan string, ch2 <-chan int) {
	select {
	case result := <-ch1:
		fmt.Println("ch1", result)
	case intResult := <-ch2:
		fmt.Println("ch2", intResult)
	}
}
