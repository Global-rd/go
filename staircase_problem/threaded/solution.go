package threaded

import (
	"fmt"
	"sync"
)

// Típusdefiníció, metódusok létrehozása lehetőségének biztosítására
type Permutations struct {
	variations [][]int
	Levels     int
	mu         sync.Mutex
}

// Típusmetódus az eredmény kiírására
func (p Permutations) Show_permutations() {
	if p.Levels == 0 {
		fmt.Println("0 lépcsőfokot nem tudok mászni!")
		return
	}
	fmt.Printf("%d lépcsőfokot %d féleképpen lehet megmászni:\n", p.Levels, len(p.variations))
	for i, perm := range p.variations {
		fmt.Printf("%d. - [%s]\n", i+1, perm)
	}
}

func (p Permutations) Calc_permutations() {
	fmt.Println("Calculating")
}
