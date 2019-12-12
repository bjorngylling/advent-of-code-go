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
	arb

	exit = 99
)

type Memory []int

type paramMode = int

const (
	modePosition paramMode = iota
	modeImmediate
	modeRelativeBase
)

type Computer struct {
	Mem     Memory
	relBase int
	In      chan int
	Out     chan int
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

func parseInstr(instr int) (operation, paramMode, paramMode, paramMode) {
	return instr % 100, instr / 100 % 10, instr / 1000 % 10, instr / 10000 % 10
}

func (c *Computer) readMem(ptr int, mode paramMode) int {
	addr := c.getAddr(mode, ptr)
	if len(c.Mem) > addr && addr >= 0 {
		return c.Mem[addr]
	} else {
		return 0
	}
}

func (c *Computer) writeMem(ptr int, mode paramMode, value int) {
	addr := c.getAddr(mode, ptr)
	if len(c.Mem) <= addr {
		c.Mem = append(c.Mem, make(Memory, addr-len(c.Mem)+50)...)
	}
	c.Mem[addr] = value
}

func (c *Computer) getAddr(mode paramMode, ptr int) int {
	switch mode {
	case modePosition:
		return c.Mem[ptr]
	case modeImmediate:
		return ptr
	case modeRelativeBase:
		return c.relBase + c.Mem[ptr]
	}
	return 0
}

func (c *Computer) Run() {
	for instrPtr := 0; c.Mem[instrPtr] != exit; {
		op, p1Mode, p2Mode, tarMode := parseInstr(c.Mem[instrPtr])
		switch op {
		case add:
			c.writeMem(instrPtr+3, tarMode, c.readMem(instrPtr+1, p1Mode)+c.readMem(instrPtr+2, p2Mode))
			instrPtr += 4
		case mul:
			c.writeMem(instrPtr+3, tarMode, c.readMem(instrPtr+1, p1Mode)*c.readMem(instrPtr+2, p2Mode))
			instrPtr += 4
		case in:
			c.writeMem(instrPtr+1, p1Mode, <-c.In)
			instrPtr += 2
		case out:
			c.Out <- c.readMem(instrPtr+1, p1Mode)
			instrPtr += 2
		case jt:
			if c.readMem(instrPtr+1, p1Mode) != 0 {
				instrPtr = c.readMem(instrPtr+2, p2Mode)
			} else {
				instrPtr += 3
			}
		case jf:
			if c.readMem(instrPtr+1, p1Mode) == 0 {
				instrPtr = c.readMem(instrPtr+2, p2Mode)
			} else {
				instrPtr += 3
			}
		case lt:
			if c.readMem(instrPtr+1, p1Mode) < c.readMem(instrPtr+2, p2Mode) {
				c.writeMem(instrPtr+3, tarMode, 1)
			} else {
				c.writeMem(instrPtr+3, tarMode, 0)
			}
			instrPtr += 4
		case eq:
			if c.readMem(instrPtr+1, p1Mode) == c.readMem(instrPtr+2, p2Mode) {
				c.writeMem(instrPtr+3, tarMode, 1)
			} else {
				c.writeMem(instrPtr+3, tarMode, 0)
			}
			instrPtr += 4
		case arb:
			c.relBase += c.readMem(instrPtr+1, p1Mode)
			instrPtr += 2
		default:
			fmt.Printf("unknown operation: ptr=%d, op=%d, p1Mode=%d, p2Mode=%d, tarMode=%d\n", instrPtr, op, p1Mode, p2Mode, tarMode)
			os.Exit(1)
		}
	}
}
