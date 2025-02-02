package staircase

import "testing"

func TestCalcWays(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
		{10, 89},
	}

	for _, test := range tests {
		if got := CalcWays(test.n); got != test.expected {
			t.Errorf("CalcWays(%d) = %d, want %d", test.n, got, test.expected)
		}
	}
}
