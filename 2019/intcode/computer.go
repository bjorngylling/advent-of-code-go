package intcode

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"os"
	"strings"
)

type operation = int

const (
	nop operation = iota
	add
	mul
	in
	out
	jt
	jf
	lt
	eq

	exit = 99
)

type Memory []int

type paramMode = int

const (
	modePosition paramMode = iota
	modeImmediate
)

type Computer struct {
	Mem Memory
	In  chan int
	Out chan int
}

func Init(program string, in chan int, out chan int) (Computer, error) {
	computer := Computer{In: in, Out: out}
	for _, instr := range strings.Split(program, ",") {
		computer.Mem = append(computer.Mem, util.GetInt(instr))
	}
	if computer.Mem == nil {
		return computer, fmt.Errorf("unable parse program from input")
	}
	return computer, nil
}

func param(reg Memory, ptr int, mode paramMode) *int {
	switch mode {
	case modePosition:
		return &reg[reg[ptr]]
	case modeImmediate:
		return &reg[ptr]
	}
	return nil
}

func parseInstr(instr int) (operation, paramMode, paramMode, paramMode) {
	return instr % 100, instr / 100 % 10, instr / 1000 % 10, instr / 10000 % 10
}

func (c Computer) Run() {
	for instrPtr := 0; c.Mem[instrPtr] != exit; {
		op, p1Mode, p2Mode, tarMode := parseInstr(c.Mem[instrPtr])
		switch op {
		case add:
			*param(c.Mem, instrPtr+3, tarMode) = *param(c.Mem, instrPtr+1, p1Mode) + *param(c.Mem, instrPtr+2, p2Mode)
			instrPtr += 4
		case mul:
			*param(c.Mem, instrPtr+3, tarMode) = *param(c.Mem, instrPtr+1, p1Mode) * *param(c.Mem, instrPtr+2, p2Mode)
			instrPtr += 4
		case in:
			*param(c.Mem, instrPtr+1, tarMode) = <-c.In
			instrPtr += 2
		case out:
			c.Out <- *param(c.Mem, instrPtr+1, p1Mode)
			instrPtr += 2
		case jt:
			if *param(c.Mem, instrPtr+1, p1Mode) != 0 {
				instrPtr = *param(c.Mem, instrPtr+2, p2Mode)
			} else {
				instrPtr += 3
			}
		case jf:
			if *param(c.Mem, instrPtr+1, p1Mode) == 0 {
				instrPtr = *param(c.Mem, instrPtr+2, p2Mode)
			} else {
				instrPtr += 3
			}
		case lt:
			if *param(c.Mem, instrPtr+1, p1Mode) < *param(c.Mem, instrPtr+2, p2Mode) {
				*param(c.Mem, instrPtr+3, tarMode) = 1
			} else {
				*param(c.Mem, instrPtr+3, tarMode) = 0
			}
			instrPtr += 4
		case eq:
			if *param(c.Mem, instrPtr+1, p1Mode) == *param(c.Mem, instrPtr+2, p2Mode) {
				*param(c.Mem, instrPtr+3, tarMode) = 1
			} else {
				*param(c.Mem, instrPtr+3, tarMode) = 0
			}
			instrPtr += 4
		default:
			fmt.Printf("unknown operation: ptr=%d, op=%d, p1Mode=%d, p2Mode=%d, tarMode=%d\n", instrPtr, op, p1Mode, p2Mode, tarMode)
			os.Exit(1)
		}
	}
}
