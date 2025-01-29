package threaded

import (
	"fmt"
	"sync"
)

// Típusdefiníció
type permutations struct {
	variations [][]int // A lépések tárolása, 2 dimenziós slice
	levels     int     // A megmászandő lépcsők száma
	mu         sync.Mutex
	wg         sync.WaitGroup
}

// Típusmetódus az eredmény kiírására
func (p permutations) Show_permutations() {
	if p.levels == 0 {
		fmt.Println("0 lépcsőfokot nem tudok mászni!")
		return
	}
	fmt.Printf("%d lépcsőfokot %d féleképpen lehet megmászni:\n", p.levels, len(p.variations))
	for i, perm := range p.variations {
		fmt.Printf("%d. - [%s]\n", i+1, perm)
	}
}

// Elértük a lépcsők tetejét?
func (p permutations) check_if_on_top(steps []int) bool {
	current_level := 0
	for _, steps := range steps {
		current_level += steps
	}
	if current_level == p.levels {
		return true
	}
	return false
}

// Túlmentünk?
func (p permutations) overdid_levels(steps []int) bool {
	current_level := 0
	for _, steps := range steps {
		current_level += steps
	}
	if current_level > p.levels {
		return true
	}
	return false
}

// A rekurzív worker
func (p permutations) calculate(variation []int) {
	if p.overdid_levels(variation) {
		p.wg.Done()
		return
	}
	if p.check_if_on_top(variation) {
		p.mu.Lock()
		p.variations = append(p.variations, variation)
		p.mu.Unlock()
		p.wg.Done()
		return
	}

	one_new_step := append(variation, 1)
	two_new_steps := append(variation, 2)
	p.wg.Add(2)
	go p.calculate(one_new_step)
	go p.calculate(two_new_steps)
	p.wg.Wait()
}

// Megoldások kiszámítása
func (p permutations) calc_permutations() {
	if p.levels == 0 {
		return
	}

	var initial []int
	var one_new_step = append(initial, 1)
	var two_new_steps = append(initial, 2)
	p.wg.Add(2)
	go p.calculate(one_new_step)
	go p.calculate(two_new_steps)

	p.wg.Wait()
}

// Inicializáslás
func NewPermutations(levels int) *permutations {
	p := new(permutations)
	p.levels = levels
	p.calc_permutations()
	return p
}
