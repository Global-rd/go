package threaded

import (
	"fmt"
)

// Egy variávió reprezentációja
type varia []int

// Struct definíció
type permutations struct {
	variations []varia // A lépések tárolása, 2 dimenziós slice
	levels     int     // A megmászandő lépcsők száma
}

// Típusmetódus az eredmény kiírására
func (p permutations) Show_permutations() {
	if p.levels == 0 {
		fmt.Println("0 lépcsőfokot nem tudok mászni!")
		return
	}
	fmt.Printf("%d lépcsőfokot %d féleképpen lehet megmászni:\n", p.levels, len(p.variations))
	for i, variation := range p.variations {
		fmt.Printf("%d. - %d\n", i+1, variation)
	}
}

// Szintjelző. ha negatív, akkor még nem értünk fel
// ha 0, akkor pont a tetjén állunk, ha pozitív, akkor
// túlmentünk
func (p *permutations) leveling(variation varia) int {
	current_level := 0
	for _, step := range variation {
		current_level += step
	}
	return current_level - p.levels
}

// A rekurzív worker
func (p *permutations) worker(variation varia) {

	// 1. meghatározzuk, hol állunk
	leveling := p.leveling(variation)

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
			// Ha túlmentünk, biztosan rossz a variáció
			case true:
				return
			// Ha túl alacsony, további variációkat generálunk, és új rekurzív workereket indítunk
			case false:
				{
					p.worker(append(variation, 1))
					p.worker(append(variation, 2))
				}
			}
		}
	}

	/*
	   // Ha túlmentünk, állítsuk le a rutint

	   	if p.above_levels(variation) {
	   		return
	   	}

	   // Ha a tetjén vagyunk, jó a megoldás, rögzítsük a variációt és állítsuk le a rutint

	   	if p.check_if_on_top(variation) {
	   		return
	   	}

	   // Ha még nem  értünk fel, két új ággal hívjuk újra a worker-t

	   	if p.below_levels(variation) {
	   		p.worker(append(variation, 1))
	   		p.worker(append(variation, 2))
	   		return
	   	}
	*/
}

// Inicializálás
func NewPermutations(levels int) *permutations {
	p := new(permutations)
	p.levels = levels
	p.variations = make([]varia, 0)

	var variation = make(varia, 0)

	p.worker(append(variation, 1))
	p.worker(append(variation, 2))

	return p
}
