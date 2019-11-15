package main

import (
	"reflect"
	"testing"
)

var input = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

func TestParseInput(t *testing.T) {
	s, rules := parseInput(input)
	expected := State{true, false, false, true, false, true, false, false, true, true, false, false, false, false, false, false, true, true, true, false, false, false, true, true, true}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected initial state to be %s but was %s", expected, s)
	}
	if len(rules) != 14 {
		t.Errorf("Expected rules to contain 14 rules but contained %d", len(rules))
	}
}

func TestStateNote(t *testing.T) {
	s, _ := parseInput(input)
	r := s.Note(0)
	expected := Note{false, false, false, false, true}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected Note(0) to be %s but was %s", expected, r)
	}
}

func TestStep(t *testing.T) {
	s, rules := parseInput(input)
	r, zI := step(s, rules)
	expected := State{false, true, false, false, false, true, false, false, false, false, true, false, false, false, false, false, true, false, false, true, false, false, true, false, false, true, false}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected state to be %s after 1 step but was %s", expected, r)
	}
	if zI != -1 {
		t.Errorf("Expected zero index diff to be -1 after step 1 but was %d", zI)
	}

	r, zI = step(r, rules)
	expected = State{false, true, true, false, false, true, true, false, false, false, true, true, false, false, false, false, true, false, false, true, false, false, true, false, false, true, true, false}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected state to be %s after 2 steps but was %s", expected, r)
	}
	if zI != 0 {
		t.Errorf("Expected zero index diff to be 0 after step 2 but was %d", zI)
	}

	r, zI = step(r, rules)
	expected = State{false, true, false, true, false, false, false, true, false, false, true, false, true, false, false, false, false, true, false, false, true, false, false, true, false, false, false, true, false}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected state to be %s after 3 steps but was %s", expected, r)
	}
	if zI != -1 {
		t.Errorf("Expected zero index diff to be -1 after step 3 but was %d", zI)
	}

	for i := 0; i < 17; i++ {
		r, zI = step(r, rules)
	}
	expected = State{false, true, false, false, false, false, true, true, false, false, false, false, true, true, true, true, true, false, false, false, true, true, true, true, true, true, true, false, false, false, false, true, false, true, false, false, true, true, false}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected state to be %s after 20 steps but was %s", expected, r)
	}
	if zI != 0 {
		t.Errorf("Expected zero index diff to be 0 after step 20 but was %d", zI)
	}
}

func TestPlantNumberSumAfterSteps(t *testing.T) {
	s, rules := parseInput(input)
	r := plantNumberSumAfterSteps(s, rules, 20)
	expected := 325
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected total number of plants to be %d after 20 generations but was %d", expected, r)
	}
}
