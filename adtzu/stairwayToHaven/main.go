package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter number of stairs: ")
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(calculateSteps(n))
}

func calculateSteps(n int) int {

	/* Imagine you are climbing a staircase with n steps. You can either:

	    	Take 1 step at a time, or
	    	Take 2 steps at a time.
			Your task is to determine the total number of distinct ways you can climb to the top of the staircase.
	*/

	if n <= 0 {
		return 0
	} else if n <= 2 {
		return n
	} else {
		return (calculateSteps(n-1) + calculateSteps(n-2))
	}
}
