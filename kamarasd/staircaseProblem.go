package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

func main() {

	os.Setenv("STEPS", "4")
	var n int

	n = readChar(n)

	if(n == 0) {
		fmt.Println("\nUser not defined strings to compare gathering default value from env!!\n")
		env := os.Getenv("STEPS")		
		fmt.Println("\nUser not defined any steps to calculate, gathering default value from env!!\n")
		m, err := strconv.Atoi(env)
		if(err != nil) {
			log.Fatal("ERROR OCCURED TERMINATE")
		}
		n = m
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
			return ((n-1) + (n-2))
	 }
}

func readChar(n int) int {

	fmt.Println("enter the number of steps: ")
	fmt.Scanf("%d", &n)

	return n
}
