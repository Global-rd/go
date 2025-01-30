package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	steps, err := readSteps()
	if err != nil {
		fmt.Println("Invalid input: ", err)
		os.Exit(1)
	}
	result := solveStaircase(steps)
	fmt.Println(result)
}

func solveStaircase(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	default:
		return solveStaircase(n-1) + solveStaircase(n-2)
	}
}

func readSteps() (int, error) {
	if stepsEnv := os.Getenv("STEPS"); stepsEnv != "" {
		steps, err := strconv.Atoi(stepsEnv)
		if err != nil {
			fmt.Println("Invalid input from environment variable:", err)
		} else {
			return steps, nil
		}
	}

	fmt.Print("Enter the number of steps: ")
	stepsScan := 0
	_, err := fmt.Scanln(&stepsScan)
	if err != nil {
		return 0, err
	}
	return stepsScan, nil
}
