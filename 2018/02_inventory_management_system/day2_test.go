package main

import "testing"

func TestGenerateChecksum(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	checksum := generateChecksum(input)
	if checksum != 12 {
		t.Errorf("Expected 12 got %d for input %v", checksum, input)
	}
}

func TestHammingDistance(t *testing.T) {
	a, b := "abcd", "abcd"
	r := hammingDistance(a, b)
	if r != 0 {
		t.Errorf("Expected 0 got %d for input %q, %q", r, a, b)
	}

	a, b = "abcd", "abce"
	r = hammingDistance(a, b)
	if r != 1 {
		t.Errorf("Expected 1 got %d for input %q, %q", r, a, b)
	}

	a, b = "abcd", "bcda"
	r = hammingDistance(a, b)
	if r != 4 {
		t.Errorf("Expected 4 got %d for input %q, %q", r, a, b)
	}
}

func TestFindSimilarIds(t *testing.T) {
	input := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	similarIds := findSimilarIds(input)
	if similarIds[0] != "fghij" && similarIds[1] != "fguij" {
		t.Errorf("Expected [fghij fghij] got %v for input %v", similarIds, input)
	}
}

func TestFindCommonId(t *testing.T) {
	input := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	commonId := findCommonId(input)
	if commonId != "fgij" {
		t.Errorf("Expected fgij got %v for input %v", commonId, input)
	}

	input = []string{"ighfbbyijnoumxjlxevacpwqtr", "ighfbsyijnoumxjlxevacpwqtr"}
	commonId = findCommonId(input)
	if commonId != "ighfbyijnoumxjlxevacpwqtr" {
		t.Errorf("Expected ighfbyijnoumxjlxevacpwqtr got %v for input %v", commonId, input)
	}
}
