package main

import (
	"fmt"
	. "github.com/bjorngylling/advent-of-code/2018/16_chronal_classification/operations"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) ([]Sample, []Instr) {
	x := strings.Split(input, "\n\n\n\n")
	var program []Instr
	for _, ln := range strings.Split(x[1], "\n") {
		program = append(program, parseInstr(ln))
	}
	return parseSamples(x[0]), program
}

func parseSamples(input string) (r []Sample) {
	for _, s := range strings.Split(input, "\n\n") {
		r = append(r, parseSample(s))
	}
	return
}

func parseSample(input string) Sample {
	split := strings.Split(input, "\n")
	return Sample{Before: parseRegisters(split[0]), Instr: parseInstr(split[1]), After: parseRegisters(split[2])}
}

func parseInstr(s string) Instr {
	f := strings.Fields(s)
	return Instruction(mustAtoi(f[0]), mustAtoi(f[1]), mustAtoi(f[2]), mustAtoi(f[3]))
}

func parseRegisters(s string) (r Registers) {
	return Registers{mustAtoi(string(s[9])), mustAtoi(string(s[12])), mustAtoi(string(s[15])), mustAtoi(string(s[18]))}
}

func mustAtoi(s string) int {
	if n, err := strconv.Atoi(s); err != nil {
		log.Fatalf("Failed to convert %q to int", s)
	} else {
		return n
	}
	return 0
}

func guessOpCode(sample Sample) (result []int) {
	for i, o := range Ops {
		if o(sample.Instr, sample.Before).Eq(sample.After) {
			result = append(result, i)
		}
	}
	return
}

func countThreeOrMorePossibleOpCodeSamples(l []Sample) (r int) {
	for _, s := range l {
		if len(guessOpCode(s)) >= 3 {
			r += 1
		}
	}
	return
}

func learnOpCodes(samples []Sample) map[int]int {
	potentials := make(map[int][]int)
	// Find all potential op code mappings by the sample pool
	for _, s := range samples {
		newPotentials := guessOpCode(s)
		if previous, ok := potentials[s.Instr.Opcode]; !ok {
			potentials[s.Instr.Opcode] = newPotentials
		} else {
			// Keep all op codes present in both current potentials and previous potentials
			var l []int
			for _, a := range previous {
				for _, b := range newPotentials {
					if a == b {
						l = append(l, b)
					}
				}
			}
			potentials[s.Instr.Opcode] = l
		}
	}

	// Reduce potentials by removing duplicates
	for range Ops {
		for k, v := range potentials {
			if len(v) == 1 {
				for ki, vi := range potentials {
					if k == ki {
						continue
					}
					if i, ok := findFirst(vi, v[0]); ok {
						potentials[ki] = append(vi[:i], vi[i+1:]...)
					}
				}
			}
		}
	}

	// Create the lookup table
	result := make(map[int]int, len(Ops))
	for k, v := range potentials {
		result[k] = v[0]
	}
	return result
}

func findFirst(l []int, n int) (int, bool) {
	for i, p := range l {
		if n == p {
			return i, true
		}
	}
	return 0, false
}

func runProgram(opCodeMapping map[int]int, program []Instr) (reg Registers) {
	for _, instr := range program {
		reg = Ops[opCodeMapping[instr.Opcode]](instr, reg)
	}
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/16_chronal_classification/day16_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	samples, program := parseInput(string(fileContent))

	fmt.Printf("Day 16 part 1 result: %+v\n", countThreeOrMorePossibleOpCodeSamples(samples))

	fmt.Printf("Day 16 part 2 result: %+v\n", runProgram(learnOpCodes(samples), program)[0])
}
