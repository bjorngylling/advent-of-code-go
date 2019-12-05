package main

import (
	"bytes"
	"fmt"
	"io"
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

	p1memory := make(memory, len(program))
	copy(p1memory, program)
	var p1 bytes.Buffer
	run(p1memory, strings.NewReader("1"), &p1)
	p2memory := make(memory, len(program))
	copy(p2memory, program)
	var p2 bytes.Buffer
	run(p2memory, strings.NewReader("5"), &p2)

	return strings.TrimLeft(p1.String(), "0"), strings.TrimLeft(p2.String(), "0")
}

func run(reg memory, reader io.Reader, writer io.Writer) memory {
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
			_, _ = fmt.Fscanf(reader, "%d", param(reg, instrPtr+1, tarMode))
			instrPtr += 2
		case OUT:
			_, _ = fmt.Fprintf(writer, "%d", *param(reg, instrPtr+1, p1Mode))
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
