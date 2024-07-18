package main

import "testing"

// TestFactorial tests the Factorial function with various input cases.
func TestFactorial(t *testing.T) {
    testCases := []struct {
        n        int
        expected int
    }{
        {0, 1},
        {1, 1},
        {2, 2},
        {3, 6},
        {4, 24},
        {5, 120},
        {-1, -1},
    }

    for _, tc := range testCases {
        result := Factorial(tc.n)
        if result != tc.expected {
            t.Errorf("Factorial(%d) = %d; expected %d", tc.n, result, tc.expected)
        }
    }
}
