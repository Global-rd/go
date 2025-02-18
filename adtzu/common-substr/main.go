package main

import (
	"fmt"
)

func main() {
	var string1 = "a4bcdefghij"
	var string2 = "1234a4bc5678"

	var longestFound string
	fmt.Println(lcs(string1, string2, longestFound, 0, 0))
}

func lcs(str1 string, str2 string, longestFound string, indexi int, indexj int) string {

	var commonStr string = ""

	for i := indexi; i < len(str1); i++ {

		for j := indexj; j < len(str2); j++ {

			if str1[i] == str2[j] && str1[i] != ' ' && str2[j] != ' ' {

				commonStr += string(str1[i])
				for i2 := i + 1; i2 < len(str1); i2++ {

					for j2 := j + 1; j2 < len(str2); j2++ {

						if i2 >= len(str1) {
							j2 = len(str2)
							continue
						}

						if str1[i2] == str2[j2] && str1[i2] != ' ' && str2[j2] != ' ' {
							commonStr += string(str1[i2])
							i2++
						} else if commonStr != "" {
							i2 = len(str1)
							j2 = len(str2)
							if len(longestFound) <= len(commonStr) {
								longestFound = commonStr
							}
							commonStr = ""
						}
					}
				}
			}

			if commonStr != "" {
				if len(longestFound) <= len(commonStr) {
					longestFound = commonStr
				}
			}
			commonStr = ""
		}
	}
	return longestFound
}
