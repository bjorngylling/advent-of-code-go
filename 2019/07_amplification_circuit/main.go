package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type instruction = int
type paramMode = int

const (
	NIL = iota
	ADD
	MUL
	IN
	OUT
	JT
	JF
	LT
	EQ

	EXIT = 99
)

type memory []instruction

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func solve(input string) (string, string) {
	var program memory
	for _, instr := range strings.Split(input, ",") {
		if i, e := strconv.Atoi(instr); e == nil {
			program = append(program, i)
		}
	}
	if program == nil {
		panic(fmt.Errorf("reg is empty, unable parse to any instructions from input"))
	}

	phaseSetting := []int{0, 1, 2, 3, 4}
	part1 := 0
	for p := make([]int, len(phaseSetting)); p[0] < len(p); nextPerm(p) {
		thrust := 0
		for _, i := range getPerm(phaseSetting, p) {
			out := make(chan int, 10)
			in := make(chan int, 10)
			in <- i
			in <- thrust
			mem := make(memory, len(program))
			copy(mem, program)
			run(mem, in, out)
			thrust = <-out
		}
		if thrust > part1 {
			part1 = thrust
		}
	}

	return strconv.Itoa(part1), ""
}

func run(reg memory, in chan int, out chan int) memory {
	for instrPtr := 0; reg[instrPtr] != EXIT; {
		op, p1Mode, p2Mode, tarMode := parseOp(reg[instrPtr])
		switch op {
		case ADD:
			*param(reg, instrPtr+3, tarMode) = *param(reg, instrPtr+1, p1Mode) + *param(reg, instrPtr+2, p2Mode)
			instrPtr += 4
		case MUL:
			*param(reg, instrPtr+3, tarMode) = *param(reg, instrPtr+1, p1Mode) * *param(reg, instrPtr+2, p2Mode)
			instrPtr += 4
		case IN:
			*param(reg, instrPtr+1, tarMode) = <-in
			instrPtr += 2
		case OUT:
			out <- *param(reg, instrPtr+1, p1Mode)
			instrPtr += 2
		case JT:
			if *param(reg, instrPtr+1, p1Mode) != 0 {
				instrPtr = *param(reg, instrPtr+2, p2Mode)
			} else {
				instrPtr += 3
			}
		case JF:
			if *param(reg, instrPtr+1, p1Mode) == 0 {
				instrPtr = *param(reg, instrPtr+2, p2Mode)
			} else {
				instrPtr += 3
			}
		case LT:
			if *param(reg, instrPtr+1, p1Mode) < *param(reg, instrPtr+2, p2Mode) {
				*param(reg, instrPtr+3, tarMode) = 1
			} else {
				*param(reg, instrPtr+3, tarMode) = 0
			}
			instrPtr += 4
		case EQ:
			if *param(reg, instrPtr+1, p1Mode) == *param(reg, instrPtr+2, p2Mode) {
				*param(reg, instrPtr+3, tarMode) = 1
			} else {
				*param(reg, instrPtr+3, tarMode) = 0
			}
			instrPtr += 4
		default:
			fmt.Printf("unknown operation: ptr=%d, op=%d, p1Mode=%d, p2Mode=%d, tarMode=%d\n", instrPtr, op, p1Mode, p2Mode, tarMode)
			os.Exit(1)
		}
	}
	return reg
}

func param(reg memory, ptr int, mode paramMode) *int {
	switch mode {
	case 0:
		return &reg[reg[ptr]]
	case 1:
		return &reg[ptr]
	}
	return nil
}

func parseOp(instr int) (instruction, paramMode, paramMode, paramMode) {
	return instr % 100, instr / 100 % 10, instr / 1000 % 10, instr / 10000 % 10
}

func (r memory) String() string {
	var strSlice []string
	for _, i := range r {
		strSlice = append(strSlice, strconv.Itoa(i))
	}
	return strings.Join(strSlice, ",")
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
