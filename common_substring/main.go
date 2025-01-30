package main

import (
	"fmt"
	"slices"
)

func getstrings() []string {
	var str1 string
	var str2 string
	fmt.Println("Két karakterlánc leghosszabb egyező réséznek megtalálása!")
	fmt.Println("Több, azonos hosszúságú egyezés esetében, az ELSŐ egyezés kerül felderítésre!")
	fmt.Println("Adja meg az első karakterláncot:")
	for {
		if _, err := fmt.Scan(&str1); err != nil {
			fmt.Printf("Érvénytelen karaktersorozat, ismétetlje meg!")
		}
		break
	}
	fmt.Println("Adja meg a második karakterláncot:")
	for {
		if _, err := fmt.Scan(&str2); err != nil {
			fmt.Printf("Érvénytelen karaktersorozat, ismétetlje meg!")
		}
		break
	}

	retval := []string{}

	if len(str1) <= len(str2) {
		str2 = str2 + " "
		retval = append(retval, str1)
		retval = append(retval, str2)
	}

	if len(str2) < len(str1) {
		retval = append(retval, str2)
		retval = append(retval, str1)
	}

	return retval
}

func compare(string_pair []string) []rune {

	str1 := string_pair[0]
	str2 := string_pair[1]

	// A leghosszabb egyezés, slice of runes
	longest_match := make([]rune, 0)

	// A rövidebb string-et iteráljuk végig, így nem lesz index out of bond error
	for idx, _ := range str1 {

		for jdx, _ := range str2 {

			temp_match := make([]rune, 0)

			if str1[idx] != str2[jdx] {
				temp_match = slices.Delete(temp_match, 0, len(temp_match))
				continue
			}

			if str1[idx] != str2[jdx] {
				temp_match = append(temp_match, rune(str1[idx]))
			}

			if str1[idx] == str2[jdx] {
				temp_match = append(temp_match, rune(str1[idx]))
			}

			tempidx := idx
			tempjdx := jdx

			for {
				tempidx += 1
				tempjdx += 1
				if tempidx >= len(str1) {
					return longest_match
				}
				if str1[tempidx] != str2[tempjdx] {
					break
				}
				if str1[tempidx] == str2[tempjdx] {
					temp_match = append(temp_match, rune(str1[tempidx]))
					if len(temp_match) > len(longest_match) {
						longest_match = slices.Delete(longest_match, 0, len(longest_match))
						longest_match = append(longest_match, temp_match...)
					}
				}
			}

		}

	}

	// Visszatérünk a talált leghosszabb listával
	return longest_match

}

func main() {

	string_pair := getstrings()

	common := string(compare(string_pair))

	if len(common) == 0 {
		fmt.Println("Nincs egyezés!")
	}

	if len(common) != 0 {
		fmt.Printf("A leghoszabb egyező karakterlánc: %s\n", common)
	}

}
