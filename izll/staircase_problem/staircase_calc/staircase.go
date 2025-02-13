package staircase_calc

// CalcWays calculates the number of ways to climb n stairs.
func CalcWays(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return CalcWays(n-1) + CalcWays(n-2)
}
