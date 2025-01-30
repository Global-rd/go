package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	stringsToCompare, err := readStringToCompare()
	if err != nil {
		fmt.Println("Error reading input: ", err)
		os.Exit(1)
	}

	longestCommonSubString := findLongestCommonSubString(stringsToCompare[0], stringsToCompare[1])
	if longestCommonSubString == "" {
		fmt.Println("No common substring found")
	} else {
		fmt.Println("Longest common substring: ", longestCommonSubString)
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

func readStringToCompare() ([]string, error) {
	if stringsToCompareEnv := os.Getenv("STRINGS_TO_COMPARE"); stringsToCompareEnv != "" {
		split := strings.Split(stringsToCompareEnv, " ")
		if len(split) != 2 {
			fmt.Println("Invalid input from environment variable: expected 2 strings separated by a space")
		} else {
			return split, nil
		}
	}

	fmt.Print("Enter strings to compare: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		if len(split) != 2 {
			return nil, errors.New("invalid input expected 2 strings separated by a space")
		}
		return split, nil
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("no input provided")
}
