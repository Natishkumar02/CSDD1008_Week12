package main

import "testing"

func TestAddition(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 4, 6},
		{2, 54, 56},
		{1, 0, 1},
		{6, -2, 4},
		{-2, -2, -4},
	}

	for _, tc := range testCases {
		result := Addition(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Addition(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
