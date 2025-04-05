package main

import "fmt"

func RunFanOut() {
	in := make(chan int)

	outa := make(chan int)
	outb := make(chan int)
	outc := make(chan int)

	go func() {
		in <- 1
		in <- 2
		close(in)
	}()

	fanOut(in, outa, outb, outc)

	select {
	case v := <-outa:
		fmt.Println("a", v)
	case v := <-outb:
		fmt.Println("b", v)
	case v := <-outc:
		fmt.Println("c", v)
	}

	// fmt.Println("finished")
}

func fanOut(in <-chan int, outa, outb, outc chan<- int) {
	go func() {
		for v := range in {
			select {
			case outa <- v:
			case outb <- v:
			case outc <- v:
			}
		}
	}()
}
