package main

import (
	"testing"
)

func TestSolveStaircase(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, test := range tests {
		result := solveStaircase(test.input)
		if result != test.output {
			t.Errorf("Expected %d, got %d", test.output, result)
		}
	}
}
