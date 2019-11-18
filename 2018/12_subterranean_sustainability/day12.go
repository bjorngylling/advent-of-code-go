package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type State []bool
type Note [5]bool
type Rules map[Note]struct{}

func (state State) String() (s string) {
	for _, b := range state {
		if b {
			s += "#"
		} else {
			s += "."
		}
	}
	return
}

func (n Note) String() (s string) {
	for _, b := range n {
		if b {
			s += "#"
		} else {
			s += "."
		}
	}
	return
}

// Returns the Note for a given position in the State, any part of the Note that is outside the state will be considered empty.
func (state State) Note(pos int) (n Note) {
	copy(n[:], append(State{false, false, false, false}, append(state, false, false, false, false)...)[pos:])
	return
}

func parseState(s string) (l State) {
	for i := 0; i < len(s); i++ {
		l = append(l, s[i] == '#')
	}
	return
}

func parseInput(input string) (state State, rules Rules) {
	rules = make(Rules)
	lines := strings.Split(input, "\n")
	state = parseState(lines[0][strings.Index(lines[0], "#"):])
	for _, ln := range lines[2:] {
		if ln[9] == '#' {
			r := [5]bool{}
			for i := 0; i < 5; i++ {
				r[i] = ln[i] == '#'
			}
			rules[r] = struct{}{}
		}
	}
	return
}

func step(s State, r Rules) (State, int) {
	ns := make(State, len(s)+4)
	zeroIndex := -3
	for i := 0; i < len(ns); i++ {
		_, ns[i] = r[s.Note(i)]
	}
	for i := 0; i < len(ns); i++ {
		if ns[i] {
			ns = ns[i-1:]
			zeroIndex += i
			break
		}
	}
	for i := len(ns) - 1; i > 0; i-- {
		if ns[i] {
			ns = ns[:i+2]
			break
		}
	}
	return ns, zeroIndex
}

func sumPlantNums(s State, zeroIndex int) (sum int) {
	for i, b := range s {
		if b {
			sum += i + zeroIndex
		}
	}
	return
}

func plantNumberSumAfterSteps(s State, r Rules, steps int) int {
	zeroIndex := 0
	for i := 0; i < steps; i++ {
		var zI int
		s, zI = step(s, r)
		zeroIndex += zI
	}
	return sumPlantNums(s, zeroIndex)
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/12_subterranean_sustainability/day12_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	state, rules := parseInput(string(fileContent))

	fmt.Printf("Day 12 part 1 result: %+v\n", plantNumberSumAfterSteps(state, rules, 20))

	gen1k := plantNumberSumAfterSteps(state, rules, 1000)
	d := gen1k - plantNumberSumAfterSteps(state, rules, 999)
	fmt.Printf("Day 12 part 2 result: %+v\n", gen1k+(50000000000-1000)*d)
}
