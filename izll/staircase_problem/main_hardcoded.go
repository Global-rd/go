package main

import (
	"fmt"
	"staircase_problem/staircase"
)

const Stairs = 10

func main() {
	calcWaysFixed()
}

func calcWaysFixed() {
	fmt.Printf("All possible ways to climb %d stairs: %d\n", Stairs, staircase.CalcWays(Stairs))
}
