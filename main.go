package main

import (
	"fmt"
	"strings"
)

func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	s := "Hello, World!"
	fmt.Printf("Number of vowels in '%s' is %d\n", s, CountVowels(s))
}
