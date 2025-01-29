package tree

import (
	"fmt"
)

// Egy variáció reprezentációja - typedef
type varia []int

// A probléma perezentáviója
type permutations struct {
	variations []varia // A lépések tárolása, 2 dimenziós slice
	levels     int     // A megmászandő lépcsők száma
}

// Típusmetódus az eredmény kiírására
func (p *permutations) Show_permutations() {
	if p.levels == 0 {
		fmt.Println("0 lépcsőfokot nem tudok mászni!")
		return
	}
	fmt.Printf("%d lépcsőfokot %d féleképpen lehet megmászni:\n", p.levels, len(p.variations))
	for i, variation := range p.variations {
		steps := 0
		for _, s := range variation {
			steps += s
		}
		fmt.Printf("%d. - %d - Összesen %d lépcsőfok.\n", i+1, variation, steps)
	}
}

func (p *permutations) append_variation(variation varia) {
	var temp = make(varia, 0)
	temp = append(temp, variation...)
	p.variations = append(p.variations, temp)
}

// Inicializálás
func NewPermutations(levels int) permutations {
	p := new(permutations)
	p.levels = levels
	p.variations = make([]varia, 0)

	variation := make(varia, 0)

	worker(append(variation, 1), p)
	worker(append(variation, 2), p)

	return *p
}

// A rekurzív worker
func worker(variation varia, p *permutations) {

	// 1. meghatározzuk, hol állunk
	current_level := 0
	for _, step := range variation {
		current_level += step
	}
	leveling := current_level - p.levels

	// 2. Ha "túlfutottunk" (nagyobb, mint 0), nem mentünk (hibás variáció)",
	// és visszatérünk a másik ágra
	if leveling > 0 {
		return
	}

	// 3. Ha pont 0, azaz felértünk a tetejére, mentjünk a variációt, és visszatérünk a másik ágra
	if leveling == 0 {
		p.append_variation(variation)
		return
	}

	// 4. Ha kevesebb, mint 0, akkor még nem értünk fel, újra hívjuk a két "lehetőséget" (ágat)
	if leveling < 0 {
		worker(append(variation, 1), p)
		worker(append(variation, 2), p)
		return
	}

}
