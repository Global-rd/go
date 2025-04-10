package utils

import (
	"math"
	"math/rand"
)

func GenerateInt() int {
	return rand.Intn(500000000)
}

func FilterFunc(num int) bool {
	return num >= 300000000
}

func CheckifPrime(num int) bool {
	if num%2 == 0 {
		return false
	}
	for i := 3; i < int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
