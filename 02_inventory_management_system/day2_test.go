package main

import "testing"

func TestGenerateChecksum(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	checksum := generateChecksum(input)
	if checksum != 12 {
		t.Errorf("Expected 12 got %d for input %v", checksum, input)
	}
}
