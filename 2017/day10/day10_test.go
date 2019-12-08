package main

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestReverseSlice(t *testing.T) {
	a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}

	if rev(a[:5]); !equal(a[:], []int{4, 3, 2, 1, 0, 5, 6, 7}) {
		t.Errorf("Expected the list to be [4 3 2 1 0 5 6 7] but was %v", a)
	}
}

func TestRotateSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	rotateR(&a, 2)
	if !equal(a, []int{4, 5, 1, 2, 3}) {
		t.Errorf("Expected the list to be [4 5 1 2 3] but was %v", a)
	}
	rotateL(&a, 1)
	if !equal(a, []int{5, 1, 2, 3, 4}) {
		t.Errorf("Expected the list to be [5 1 2 3 4] but was %v", a)
	}
}

func TestReversePartOfList(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if r := reverse(a, 0, 5); !equal(r, []int{4, 3, 2, 1, 0, 5, 6, 7, 8, 9}) {
		t.Errorf("Expected list to be [4 3 2 1 0 5 6 7 8 9] but was %v", r)
	}

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if r := reverse(a, 8, 4); !equal(r, []int{9, 8, 2, 3, 4, 5, 6, 7, 1, 0}) {
		t.Errorf("Expected list to be [9 8 2 3 4 5 6 7 1 0] but was %v", r)
	}
	if r := reverse(a, 5, 8); !equal(r, []int{7, 6, 5, 3, 4, 2, 1, 0, 9, 8}) {
		t.Errorf("Expected list to be [7 6 5 3 4 2 1 0 9 8] but was %v", r)
	}
}

func TestPart1(t *testing.T) {
	if ans := part1(readInput("day10_input")); ans != 11413 {
		t.Errorf("Expected answer for part1 to be 11413 but was %d", ans)
	}
}
