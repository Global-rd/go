package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const EnvVarNumOfSteps = "NUMBER_OF_STEPS"

func ValidateInput(numOfSteps int64) (bool, error) {
	if numOfSteps <= 0 {
		return false, errors.New("number of steps must be greater than 0")
	}
	return true, nil
}

func CalculateWaysOfClimbing(numOfSteps int) int {
	if numOfSteps <= 3 {
		return numOfSteps
	}
	oneStepLessWays := 3
	twoStepsLessWays := 2
	possibleWays := 0
	for i := 4; i <= numOfSteps; i++ {
		possibleWays = oneStepLessWays + twoStepsLessWays
		twoStepsLessWays = oneStepLessWays
		oneStepLessWays = possibleWays
	}
	return possibleWays
}

func CheckForEnvironmentVariable() (int64, bool) {
	numOfStepsStr := os.Getenv(EnvVarNumOfSteps)
	if numOfStepsStr == "" {
		return 0, false
	}
	numOfSteps, err := strconv.ParseInt(numOfStepsStr, 10, 64)
	if err != nil {
		panic(fmt.Errorf("error parsing number of steps from environment variable (%s): %w", EnvVarNumOfSteps, err))
	}
	return numOfSteps, true
}

func GetNumberOfStepsFromUserInput() int64 {
	var numOfSteps int64
	fmt.Print("Enter the number of steps: ")
	_, err := fmt.Scanln(&numOfSteps)
	if err != nil {
		panic(fmt.Errorf("error reading number of steps from user input: %w", err))
	}
	return numOfSteps
}

func main() {
	numOfSteps, ok := CheckForEnvironmentVariable()
	if !ok {
		numOfSteps = GetNumberOfStepsFromUserInput()
	}
	isValid, err := ValidateInput(numOfSteps)
	if err != nil {
		panic(err)
	}
	if isValid {
		waysOfClimbing := CalculateWaysOfClimbing(int(numOfSteps))
		fmt.Printf("There are %d ways to climb %d steps\n", waysOfClimbing, numOfSteps)
	}
}
