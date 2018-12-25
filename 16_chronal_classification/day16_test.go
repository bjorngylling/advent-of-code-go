package main

import (
	. "github.com/bjorngylling/advent-of-code-2018/16_chronal_classification/operations"
	"reflect"
	"testing"
)

func TestParseSample(t *testing.T) {
	input := `Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`
	result := parseSample(input)
	expected := Sample{Registers{3, 2, 1, 1}, Instr{9, 2, 1, 2}, Registers{3, 2, 2, 1}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}

func TestGuessOpCode(t *testing.T) {
	input := `Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`
	possibleOps := guessOpCode(parseSample(input))
	expected := []int{1, 2, 8}
	if !reflect.DeepEqual(possibleOps, expected) {
		t.Errorf("Expected %+v but was %+v", expected, possibleOps)
	}
}
