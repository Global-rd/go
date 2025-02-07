package main

import (
	"fmt"
	"slices"
)

// Struct declaration to store the comparables
type Tuple struct {
	shortest, longest string
}

// Show user interface intro
func PrintIntro() {
	fmt.Println(" _______________________________________________________________________________________ ")
	fmt.Println("|                                                                                       |")
	fmt.Println("|               Simple solution for find the longest common substring in                |")
	fmt.Println("|               an user defined substring.                                              |")
	fmt.Println("|                                                                                       |")
	fmt.Println("|            *NOTICE:                                                                   |")
	fmt.Println("|            The solution finds only the FIRST longest substring, even if there         |")
	fmt.Println("|            are more different matches!                                                |")
	fmt.Println("|                                                                                       |")
	fmt.Println("|               The solution accepts user input as input parameters. Whitespaces        |")
	fmt.Println("|               invalid parameters. The user can quit on input only \"q\" or \"Q\".         |")
	fmt.Println("|_______________________________________________________________________________________|")
}

// Get user input for strings to compare
func GetStrings() (Tuple, int) {
	// Instantiate struct to return
	strings := Tuple{}

	// Declare temporary variables to hold the user inputs
	var str1 string
	var str2 string

	for {
		// User instruction
		fmt.Println("Enter the first string! (\"q\" or \"Q\" to quit):")
		// In case of error, retry user input
		if _, err := fmt.Scan(&str1); err != nil {
			fmt.Printf("Invalid input! Try again or quit!")
			continue
		}
		// In case of user uit, return a negative integer to caller
		if str1 == "q" || str1 == "Q" {
			return strings, -1
		}
		break
	}

	for {
		// User instruction
		fmt.Println("Enter the second string! (\"q\" or \"Q\" to quit):")
		// In case of error, retry user input
		if _, err := fmt.Scan(&str2); err != nil {
			fmt.Printf("Invalid input! Try again or quit!")
			continue
		}
		// In case of user uit, return a negative integer to caller
		if str2 == "q" || str2 == "Q" {
			return strings, -1
		}

		break
	}

	// Sort user inputs by length
	// In case of equality,  add an extra whitespace to the second input
	if len(str1) == len(str2) {
		strings.shortest = str1
		strings.longest = str2 + " "
	}

	if len(str1) < len(str2) {
		strings.shortest = str1
		strings.longest = str2
	}

	if len(str2) < len(str1) {
		strings.shortest = str2
		strings.longest = str1
	}

	// return the user input in the defined struct and a positive integer
	return strings, 1
}

// Business logic
func FindLongestCommon(strings Tuple) string {
	// Temporary slice of runes to store the longest match
	longest := []rune{}
	// Iterate over the shortest string
	for i, _ := range strings.shortest {
		// Iterate iver the longest string
		for j, _ := range strings.longest {

			// Temporary variables
			temp_match := []rune{0}
			temp_idx := i
			temp_jdx := j

			for {
				// If overrun any of the strings, to prevent error, break the loop
				if temp_idx > len(strings.shortest)-1 || temp_jdx > len(strings.longest)-1 {
					break
				}
				// If the current characters doesn't match, break the loop
				if strings.shortest[temp_idx] != strings.longest[temp_jdx] {
					break
				}
				// If current characters (runes) matches, append current char to the temporary
				// container
				if strings.shortest[temp_idx] == strings.longest[temp_jdx] {
					temp_match = append(temp_match, []rune(strings.shortest)[temp_idx])
				}
				// Increment temporary counters
				temp_idx += 1
				temp_jdx += 1
			}

			// Check if the temporary contained match is longet as the current longest match
			// If so, overwrite it
			if len(temp_match) > len(longest) {
				longest = slices.Delete(longest, 0, len(longest))
				longest = append(longest, temp_match...)
			}

		}

	}
	// Return the found longest match
	return string(longest)
}

// Entry point
func main() {

	// Instruct user
	PrintIntro()

	// Repeat logic until user quit
	for {
		// Get user inputs
		strings, status := GetStrings()
		// In case of user want too quit, break the loop
		if status < 0 {
			break
		}
		// Calculate and show the first longest common substring to the user
		// The loop restarts
		longest_common := FindLongestCommon(strings)
		fmt.Println("The longest common string: ", longest_common)
	}
	// Exit program
	fmt.Println("Bye!")
}
