package main

import (
	"fmt"
	"slices"
)

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
func GetStrings() (string, string, int) {

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
			return str1, str2, -1
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
			return str1, str2, -1
		}

		break
	}

	// Sort user inputs by length
	// In case of equality,  add an extra whitespace to the second input
	if len(str1) == len(str2) {
		return str1, str2 + "", 1
	}

	if len(str1) < len(str2) {
		return str1, str2, 1
	}

	if len(str2) < len(str1) {
		return str2, str1, 1
	}

	return str1, str2, 1

}

// Business logic
func FindLongestCommon(shortest_string, longest_string string) string {
	// Temporary slice of runes to store the longest match
	longest := []rune{}
	// Iterate over the shortest string
	for i, _ := range shortest_string {
		// Iterate iver the longest string
		for j, _ := range longest_string {

			// Temporary variables
			temp_match := []rune{0}
			temp_idx := i
			temp_jdx := j

			for {
				// If overrun any of the strings, to prevent error, break the loop
				if temp_idx > len(shortest_string)-1 || temp_jdx > len(longest_string)-1 {
					break
				}
				// If the current characters doesn't match, break the loop
				if shortest_string[temp_idx] != longest_string[temp_jdx] {
					break
				}
				// If current characters (runes) matches, append current char to the temporary
				// container
				if shortest_string[temp_idx] == longest_string[temp_jdx] {
					temp_match = append(temp_match, []rune(shortest_string)[temp_idx])
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
		shortest_string, longest_string, status := GetStrings()
		// In case of user want too quit, break the loop
		if status < 0 {
			break
		}
		// Calculate and show the first longest common substring to the user
		// The loop restarts
		longest_common := FindLongestCommon(shortest_string, longest_string)
		fmt.Println("The longest common string: ", longest_common)
	}
	// Exit program
	fmt.Println("Bye!")
}
