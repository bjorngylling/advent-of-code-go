package main

import "testing"

func TestCalculateFrequency(t *testing.T) {
	input := []string{"+1", "+1", "+1"}
	freq := calculateFrequency(input)
	if freq != 3 {
		t.Errorf("Expected 3 got %d for input %v", freq, input)
	}

	input = []string{"+1", "+1", "-2"}
	freq = calculateFrequency(input)
	if freq != 0 {
		t.Errorf("Expected 0 got %d for input %v", freq, input)
	}

	input = []string{"-1", "-2", "-3"}
	freq = calculateFrequency(input)
	if freq != -6 {
		t.Errorf("Expected -6 got %d for input %v", freq, input)
	}
}

func TestFindFirstDuplicateFrequency(t *testing.T) {
	input := []string{"+1", "-1"}
	freq := findFirstDuplicateFrequency(input)
	if freq != 0 {
		t.Errorf("Expected 0 got %d for input %v", freq, input)
	}

	input = []string{"+3", "+3", "+4", "-2", "-4"}
	freq = findFirstDuplicateFrequency(input)
	if freq != 10 {
		t.Errorf("Expected 10 got %d for input %v", freq, input)
	}

	input = []string{"-6", "+3", "+8", "+5", "-6"}
	freq = findFirstDuplicateFrequency(input)
	if freq != 5 {
		t.Errorf("Expected 5 got %d for input %v", freq, input)
	}

	input = []string{"+7", "+7", "-2", "-7", "-4"}
	freq = findFirstDuplicateFrequency(input)
	if freq != 14 {
		t.Errorf("Expected 14 got %d for input %v", freq, input)
	}
}
