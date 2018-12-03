package main

import "testing"

func TestGenerateChecksum(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	result := testFunc(input)
	if result != 12 {
		t.Errorf("Expected 12 got %d for input %v", result, input)
	}
}
