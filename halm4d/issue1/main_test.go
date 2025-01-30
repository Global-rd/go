package main

import "testing"

func TestFindLongestCommonSubString(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		output string
	}{
		{"hello", "world", "l"},
		{"hello", "hello", "hello"},
		{"hello", "hell", "hell"},
		{"hello", "ello", "ello"},
		{"hello", "ell", "ell"},
		{"hello", "llo", "llo"},
	}

	for _, test := range tests {
		result := findLongestCommonSubString(test.a, test.b)
		if result != test.output {
			t.Errorf("Expected %s, got %s", test.output, result)
		}
	}
}

func TestFindSubString(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		i      int
		j      int
		output string
	}{
		{"hello", "world", 0, 0, ""},
		{"hello", "hello", 0, 0, "hello"},
		{"hello", "hello", 1, 1, "ello"},
		{"hello", "world", 2, 3, "l"},
	}

	for _, test := range tests {
		result := findSubString(test.a, test.b, test.i, test.j)
		if result != test.output {
			t.Errorf("Expected %s, got %s", test.output, result)
		}
	}
}
