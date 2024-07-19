package main

import (
	"testing"
)

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"Hello, World!", 3},
		{"", 0},
		{"AEIOU", 5},
		{"Go Programming", 4},
		{"Count the vowels!", 5},
		{"This is a test.", 4},
	}

	for _, tc := range testCases {
		result := CountVowels(tc.input)
		if result != tc.expected {
			t.Errorf("CountVowels(%q) = %d; expected %d", tc.input, result, tc.expected)
		}
	}
}
