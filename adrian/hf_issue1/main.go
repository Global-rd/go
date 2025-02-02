package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	EnvVarStr1 = "STRING_1"
	EnvVarStr2 = "STRING_2"
)

var (
	WhiteSpaces = []rune{' ', '\t', '\n', '\r'}
)

func IsWhiteSpace(r rune) bool {
	for _, w := range WhiteSpaces {
		if r == w {
			return true
		}
	}
	return false
}

func CheckForEnvironmentVariables() (string, string, error) {
	str1 := os.Getenv(EnvVarStr1)
	str2 := os.Getenv(EnvVarStr2)
	if str1 == "" || str2 == "" {
		return "", "", errors.New("missing environment variables")
	}
	return str1, str2, nil
}

func GetUserInput() (string, string, error) {
	var str1, str2 string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first string to compare: ")
	str1, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	fmt.Print("Enter the second string to compare: ")
	str2, err = reader.ReadString('\n')
	return strings.TrimSuffix(str1, "\n"), strings.TrimSuffix(str2, "\n"), nil
}

// FindCommonSubstring Implementing common substring search using Dynamic Programming (DP) Algorithm
// A two-dimensional array stores the length of common substrings
func FindCommonSubstring(str1, str2 string) (int, string) {
	lenS1, lenS2 := len(str1), len(str2)
	resultArray := make([][]int, lenS1)
	for i := range resultArray {
		resultArray[i] = make([]int, lenS2)
	}
	maxLen, startIdx, endIdx := 0, 0, 0
	for i, s1Rune := range str1 {
		for j, s2Rune := range str2 {
			if s1Rune != s2Rune || IsWhiteSpace(s1Rune) {
				resultArray[i][j] = 0
			} else {
				prevI := max(0, i-1)
				prevJ := max(0, j-1)
				resultArray[i][j] = resultArray[prevI][prevJ] + 1
				if resultArray[i][j] > maxLen {
					maxLen = resultArray[i][j]
					endIdx = i + 1
					startIdx = endIdx - maxLen
				}
			}

		}
	}
	foundSubStr := ""
	if maxLen > 0 {
		foundSubStr = str1[startIdx:endIdx]
	}
	return maxLen, foundSubStr
}

func main() {
	str1, str2, err := CheckForEnvironmentVariables()
	if err != nil {
		str1, str2, err = GetUserInput()
		if err != nil {
			fmt.Printf("Failed to read user input due to error: %v\n", err)
			os.Exit(1)
		}
	}
	if len(str1) == 0 || len(str2) == 0 {
		fmt.Println("No common substring found")
		os.Exit(1)
	}
	fmt.Printf("Comparing strings '%s' and '%s'\n", str1, str2)
	subStrlen, subStr := FindCommonSubstring(str1, str2)
	fmt.Printf("\n >>> Found longest common substring: '%s' with length %d <<<\n", subStr, subStrlen)
}
