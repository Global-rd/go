package staircase_env

import (
	"fmt"
	"os"
	"staircase_problem/staircase_calc"
	"strconv"
)

const StaircaseEnv = "STAIRCASE_COUNT"

func Run() {
	calcWaysEnv()
}

func calcWaysEnv() {
	if n, err := stairsCount(); err == nil {
		fmt.Printf("All possible ways to climb %d stairs: %d\n", n, staircase_calc.CalcWays(n))
	} else {
		fmt.Println(err)
	}
}

func stairsCount() (int, error) {
	if value, exists := os.LookupEnv(StaircaseEnv); exists {
		n, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return int(n), nil
		}
	}
	return 0, fmt.Errorf("Environment variable %s is not set or invalid", StaircaseEnv)
}
