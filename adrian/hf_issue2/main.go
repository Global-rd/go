package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const EnvVarNumOfSteps = "NUMBER_OF_STEPS"

func ValidateInput(numOfSteps int64) error {
	if numOfSteps <= 0 {
		return errors.New("number of steps must be greater than 0")
	}
	return nil
}

func CalculateWaysOfClimbing(numOfSteps int) string {
	// Up to 3 steps we have the same number of possibilities as number of steps, so just return it and we are done
	if numOfSteps <= 3 {
		return strconv.Itoa(numOfSteps)
	}

	// Otherwise let's do the calculation - each iteration should calculate the result as the sum of 2 previous results
	// Also we are saving the last 2 results for next iteration
	oneStepLessWays := big.NewInt(3)
	twoStepsLessWays := big.NewInt(2)
	possibleWays := big.NewInt(0)

	for i := 4; i <= numOfSteps; i++ {
		possibleWays.Add(oneStepLessWays, twoStepsLessWays)
		twoStepsLessWays.Set(oneStepLessWays)
		oneStepLessWays.Set(possibleWays)
	}
	return possibleWays.String()
}

func CheckForEnvironmentVariable() (int64, error) {
	numOfStepsStr := os.Getenv(EnvVarNumOfSteps)
	if numOfStepsStr == "" {
		return 0, nil
	}
	numOfSteps, err := strconv.ParseInt(numOfStepsStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return numOfSteps, nil
}

func GetNumberOfStepsFromUserInput() (int64, error) {
	var numOfSteps int64
	fmt.Print("Enter the number of steps: ")
	_, err := fmt.Scanln(&numOfSteps)
	if err != nil {
		return -1, err
	}
	return numOfSteps, nil
}

func main() {
	numOfSteps, err := CheckForEnvironmentVariable()
	if err != nil {
		fmt.Printf("Error reading number of steps from environment variable: %s.\n Please check and fix it!\n", err)
		os.Exit(1)
	}
	if numOfSteps <= 0 {
		numOfSteps, err = GetNumberOfStepsFromUserInput()
		if err != nil {
			fmt.Printf("Error reading number of steps from user input: %s.\n", err)
			os.Exit(1)
		}
	}

	err = ValidateInput(numOfSteps)
	if err != nil {
		fmt.Printf("Error validating number of steps: %s.\n", err)
		os.Exit(1)
	}

	waysOfClimbing := CalculateWaysOfClimbing(int(numOfSteps))
	fmt.Printf("There are %s ways to climb %d steps\n", waysOfClimbing, numOfSteps)

}
