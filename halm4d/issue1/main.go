package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	firstString, secondString, err := readStringToCompare()
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	longestCommonSubString := findLongestCommonSubString(firstString, secondString)
	if longestCommonSubString == "" {
		fmt.Println("No common substring found")
	} else {
		fmt.Println("Longest common substring:", longestCommonSubString)
	}

}

func findLongestCommonSubString(a, b string) string {
	longestCommonSubString := ""
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				subString := findSubString(a, b, i, j)
				if len(subString) > len(longestCommonSubString) {
					longestCommonSubString = subString
				}
			}
		}
	}
	return longestCommonSubString
}

func findSubString(a, b string, i, j int) string {
	subString := ""
	for k := 0; i+k < len(a) && j+k < len(b); k++ {
		if a[i+k] != b[j+k] {
			break
		}
		subString += string(a[i+k])
	}
	return subString
}

func readStringToCompare() (string, string, error) {
	var firstStringToCompare string
	var secondStringToCompare string

	if stringsToCompareEnv := os.Getenv("FIRST_STRINGS_TO_COMPARE"); stringsToCompareEnv != "" {
		firstStringToCompare = stringsToCompareEnv
	}
	if stringsToCompareEnv := os.Getenv("SECOND_STRINGS_TO_COMPARE"); stringsToCompareEnv != "" {
		secondStringToCompare = stringsToCompareEnv
	}

	if firstStringToCompare != "" && secondStringToCompare != "" {
		return firstStringToCompare, secondStringToCompare, nil
	}

	fmt.Print("Enter first strings to compare: ")
	if _, err := fmt.Scanln(&firstStringToCompare); err != nil {
		return "", "", err
	}
	fmt.Print("Enter second strings to compare: ")
	if _, err := fmt.Scanln(&secondStringToCompare); err != nil {
		return "", "", err
	}

	if firstStringToCompare == "" || secondStringToCompare == "" {
		return "", "", errors.New("no input provided")
	}
	return firstStringToCompare, secondStringToCompare, nil
}
