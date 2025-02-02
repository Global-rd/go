package main

import (
	"os"
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

func TestReadStepsWithEnv(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"1", 1},
		{"2", 2},
	}

	for _, test := range tests {
		_ = os.Setenv("STEPS", test.input)
		result, err := readSteps()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != test.output {
			t.Errorf("Expected %d, got %d", test.output, result)
		}
	}
}
