package main

import (
	"fmt"
	"strconv"
)

// Típusdefiníció, metódusok létrehozása lehetőségének biztosítására
type Permutations []string

// Típusmetódus az eredmény kiírására
func (p Permutations) show_permutations(stairs int) {
	if stairs == 0 {
		fmt.Println("0 lépcsőfokot nem tudok mászni!")
		return
	}
	fmt.Printf("%d lépcsőfokot %d féleképpen lehet megmászni:\n", stairs, len(p))
	for i, perm := range p {
		fmt.Printf("%d. - [%s]\n", i+1, perm)
	}
}

// A kezdőfeltételen bekérése és parse-olása
func scan_input() int {
	var input string

	// Bemenet étrvényességének ellenőrzése
	if _, err := fmt.Scan(&input); err != nil {
		fmt.Printf("Érvénytelen válasz! (q - kilépés)\n")
		return scan_input()
	}

	// kilépő karakter ellenőrzése
	if input == "q" {
		return 0
	}

	// Ellenőrzés, hogy a bemenet positív egész szám-e
	if stairs, err := strconv.Atoi(input); err == nil && stairs > 0 {
		return stairs
	}

	//
	fmt.Printf("Érvénytelen válasz! (q - kilépés)\n")
	return scan_input()
}

// Intro, hogy a user tudja, mi a feladata
func intro() int {

	fmt.Println("Add meg, hány lépcsőfokot másszak, én pedig kiszámolom, hogy 1 vagy 2 lépcsőfokos\nlépésekkel hányféleképpen tudom megmászni!")

	return scan_input()
}

func calc_permutations(permutations *Permutations, stairs int) {
	fmt.Println("Calculating")
}

func main() {
	p := new(Permutations)
	stairs := intro()
	calc_permutations(p, stairs)
	p.show_permutations(stairs)
}
