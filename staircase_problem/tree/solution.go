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
		fmt.Printf("%d. - %d\n", i+1, variation)
	}
}

// Inicializálás
func NewPermutations(levels int) *permutations {
	p := new(permutations)
	p.levels = levels
	p.variations = make([]varia, 0)

	var variation = make(varia, 0)

	worker(append(variation, 1), p)
	worker(append(variation, 2), p)

	return p
}

// A rekurzív worker
func worker(variation varia, p *permutations) {

	// 1. meghatározzuk, hol állunk
	current_level := 0
	for _, step := range variation {
		current_level += step
	}
	leveling := current_level - p.levels

	switch leveling == 0 {

	// Ha éppen a tetején, akkor rögzítjük az aktuális variációt és kilépünk
	case true:
		{
			p.variations = append(p.variations, variation)
			return
		}
	case false:
		// Ha nem pont a tetején tovább vizsgálódunk
		{
			switch leveling > 0 {
			// Ha túlmentünk, biztosan rossz a variáció, kilépünk
			case true:
				return
			// Ha túl alacsony, további variációkat generálunk, és új rekurzív worker-eket indítunk
			case false:
				{
					worker(append(variation, 1), p)
					worker(append(variation, 2), p)
					return
				}
			}
		}
	}

}
