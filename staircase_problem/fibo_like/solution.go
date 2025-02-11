package fibolike

func Fibonacci_like(stairs int) int {
	if stairs == 0 || stairs == 1 {
		return 1
	}
	return Fibonacci_like(stairs-1) + Fibonacci_like((stairs - 2))
}
