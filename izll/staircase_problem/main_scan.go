package main

import (
	"fmt"
	"staircase_problem/staircase"
	"strconv"
)

func main() {
	calcWaysInput()
}

func calcWaysInput() {
	stairs := readStairs()
	fmt.Println("All possible ways to climb the staircase:", staircase.CalcWays(stairs))
}

func readStairs() int {
	var input string
	fmt.Print("Enter your input: ")
	fmt.Scanln(&input)
	n, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return readStairs()
	}
	if n < 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return readStairs()
	}
	fmt.Println("You entered:", n)
	return int(n)
}
