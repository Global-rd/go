package main

import (
	"fmt"
)

func main() {

	var defaultSteps int = 4
	var n int

	n = readChar(n)

	if n == 0 {
		fmt.Println("User not defined any steps to calculate, gathering default value from!!")
		n = defaultSteps
	}

	fmt.Println("Steps :", n, " Total number of ways: ", calc(n))
}

func calc(n int) int {

	switch {
	case n < 1:
		return 1
	case n <= 2:
		return 2
	default:
		return (calc(n-1) + calc(n-2))
	}
}

func readChar(n int) int {

	fmt.Println("Enter the number of steps: ")
	fmt.Scanf("%d", &n)

	return n
}
