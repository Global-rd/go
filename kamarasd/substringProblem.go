package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	os.Setenv("STRING_A", "asdfghjzu")
	os.Setenv("STRING_B", "qwefghjtz")

	var str1 string
	var str2 string
	var msg string = ""

	str1, str2 = readChars(str1, str2)
	
	if(str1 != "" || str2 != "") {
		msg = "User defined two strings to get common chars."
	} else {
		msg = "User not defined strings to compare gathering default values from env!!"
		str1 = os.Getenv("STRING_A")
		str2 = os.Getenv("STRING_B")
	}

	calculateCommonChars(str1, str2, msg)
	
}

func calculateCommonChars(str1 string, str2 string, msg string) {

	fmt.Print(msg, "\nA string: ", str1, "\nB string: ", str2, "\n\n")

	var maxLength = 0
	var end = 0

	matrix := make([][]int, len(str1))

	for i := range matrix {
		matrix[i] = make([]int, len(str2))
	}

	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			if(string(str1[i - 1]) == " " || string(str2[j - 1]) == " ") {
				log.Fatal("Space found in string. NO SPACE allowed. TERMINATE")
				return
			}
			if(str1[i - 1] == str2[j - 1]) {
				matrix[i][j] = matrix[i-1][j-1] + 1
				if(matrix[i][j] > maxLength) {
					maxLength = matrix[i][j]
					end = i - 1
				}
			}
		}
	}

/*	 a s d f g h j z u
q   [0 0 0 0 0 0 0 0 0] 
w	[0 0 0 0 0 0 0 0 0]  
e	[0 0 0 0 0 0 0 0 0] 
f	[0 0 0 0 0 0 0 0 0] 
g	[0 0 0 0 1 0 0 0 0]
h	[0 0 0 0 0 2 0 0 0] 
j	[0 0 0 0 0 0 3 0 0] 
t	[0 0 0 0 0 0 0 4 0] 
z	[0 0 0 0 0 0 0 0 0] */
	fmt.Println("The common chars and length are: ")
	fmt.Println(str1[end-maxLength+1 : end+1], maxLength)
}

func readChars(str1 string, str2 string) (string, string) {

	fmt.Println("Enter string A: ")
	fmt.Scanln(&str1)
	fmt.Println("Enter string B: ")
	fmt.Scanln(&str2)

	return str1, str2
} 