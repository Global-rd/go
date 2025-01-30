package main

import (
	"fmt"
	fibolike "main/fibo_like"
	"main/tree"
	"strconv"
	"time"
)

func Benchmark(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s time cost: %s\n", name, elapsed)
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

func main() {
	stairs := intro()
	if stairs == 0 {
		return
	}
	fmt.Println("")

	// A variációk prezentálása rekurzív algoritmussal
	fmt.Println("Egyszerű algoritmus:")
	var now = time.Now()
	fmt.Printf("A variációk száma %d lépcsőfokra: %d.\n", stairs, fibolike.Fibonacci_like(stairs))
	Benchmark(now, "Egyszerű megoldás")

	fmt.Println("")
	// A variációk prezentálása tree-like feldolgozással
	// TODO fix bugs (number of variations passes, invalid variations)

	fmt.Println("Variációk listázásával:")
	now = time.Now()
	p := tree.NewPermutations(stairs)
	elapsed := time.Since(now)
	p.Show_permutations()
	fmt.Printf("Bonyolult megoldás time cost: %s\n", elapsed)

}
