package main

import "fmt"

func RunFanIn() {
	out := make(chan int)

	ina := make(chan int)
	inb := make(chan int)
	inc := make(chan int)

	go func() {
		ina <- 1
	}()

	go func() {
		inb <- 2
	}()

	go func() {
		inc <- 3
	}()

	fanIn(ina, inb, inc, out)

	fmt.Println("v", <-out)
	fmt.Println("v", <-out)
	fmt.Println("v", <-out)

	fmt.Println("finished")
}

func fanIn(ina, inb, inc <-chan int, out chan<- int) {
	go func() {
		for {
			select {
			case v := <-ina:
				out <- v
			case v := <-inb:
				out <- v
			case v := <-inc:
				out <- v
			}
		}
	}()
}
