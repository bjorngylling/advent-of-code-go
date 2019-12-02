package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type instruction = int

const (
	ADD = 1
	MUL = 2

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
	part1 := make([]instruction, len(reg))
	copy(part1, reg)
	part1[1] = 12
	part1[2] = 2
	run(part1)

	part2 := make([]instruction, len(reg))
outer:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(part2, reg)
			part2[1] = noun
			part2[2] = verb
			run(part2)
			if part2[0] == 19690720 {
				break outer
			}
		}
	}

	return strconv.Itoa(part1[0]), fmt.Sprintf("%d%d", part2[1], part2[2])
}

func run(reg memory) memory {
	for instrPntr := 0; reg[instrPntr] != EXIT; instrPntr += 4 {
		op := reg[instrPntr]
		p1Addr := reg[instrPntr+1]
		p2Addr := reg[instrPntr+2]
		tarAddr := reg[instrPntr+3]
		switch op {
		case ADD:
			reg[tarAddr] = reg[p1Addr] + reg[p2Addr]
		case MUL:
			reg[tarAddr] = reg[p1Addr] * reg[p2Addr]
		}
	}
	return reg
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
