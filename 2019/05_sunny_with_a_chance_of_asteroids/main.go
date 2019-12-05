package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type instruction = int
type paramMode = int

const (
	ADD = 1
	MUL = 2
	IN  = 3
	OUT = 4

	EXIT = 99
)

type memory []instruction

func solve(input string) (string, string) {
	var reg memory
	for _, instr := range strings.Split(input, ",") {
		if i, e := strconv.Atoi(instr); e == nil {
			reg = append(reg, i)
		}
	}
	if reg == nil {
		panic(fmt.Errorf("reg is empty, unable parse to any instructions from input"))
	}

	return "", ""
}

func run(reg memory) memory {
	for instrPntr := 0; reg[instrPntr] != EXIT; {
		op, p1Mode, p2Mode, tarMode := parseOp(reg[instrPntr])
		switch op {
		case ADD:
			*param(reg, instrPntr+3, tarMode) = *param(reg, instrPntr+1, p1Mode) + *param(reg, instrPntr+2, p2Mode)
			instrPntr += 4
		case MUL:
			*param(reg, instrPntr+3, tarMode) = *param(reg, instrPntr+1, p1Mode) * *param(reg, instrPntr+2, p2Mode)
			instrPntr += 4
		case IN:
			fmt.Scanf("%d", *param(reg, instrPntr+1, tarMode))
			instrPntr += 2
		case OUT:
			fmt.Printf("ptr=%d, output=%d\n", instrPntr, *param(reg, instrPntr+1, tarMode))
			instrPntr += 2
		}
	}
	return reg
}

func param(reg memory, pntr int, mode paramMode) *int {
	switch mode {
	case 0:
		return &reg[reg[pntr]]
	case 1:
		return &reg[pntr]
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
	run(memory{3, 0, 4, 0, 99})

	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
