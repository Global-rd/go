package staircase_hardcoded

import (
	"fmt"
	"staircase_problem/staircase_calc"
)

const Stairs = 10

func Run() {
	calcWaysFixed()
}

func calcWaysFixed() {
	fmt.Printf("All possible ways to climb %d stairs: %d\n", Stairs, staircase_calc.CalcWays(Stairs))
}
