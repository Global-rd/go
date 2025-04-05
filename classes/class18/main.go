package main

import "fmt"

func main() {
	sum := sum(add(add(generator())))

	// for v := range ch {
	// 	fmt.Println("value", v)
	// }

	fmt.Println("finished", sum)
}

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		for v := range 10 {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func add(ch <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for v := range ch {
			result <- v + 10
		}
		close(result)
	}()

	return result
}

func sum(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}

	return sum
}
