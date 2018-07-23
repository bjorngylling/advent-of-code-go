package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type instrOperation func(int, int) int

var instrOperations = map[string]instrOperation{
	"inc": func(l int, r int) int { return l + r },
	"dec": func(l int, r int) int { return l - r },
}

type boolOperation func(int, int) bool

var boolOperations = map[string]boolOperation{
	"==": func(l int, r int) bool { return l == r },
	"!=": func(l int, r int) bool { return l != r },
	"<":  func(l int, r int) bool { return l < r },
	"<=": func(l int, r int) bool { return l <= r },
	">":  func(l int, r int) bool { return l > r },
	">=": func(l int, r int) bool { return l >= r },
}

const (
	instrReg = iota
	instrOperator
	instrVal

	_ // ignore 'if' keyword
	condReg
	condOperator
	condVal
)

type instr struct {
	Reg  string
	Op   instrOperation
	Val  int
	Cond cond
}
type cond struct {
	Reg string
	Op  boolOperation
	Val int
}

type cpu struct {
	Registers map[string]int
}

func newCPU() *cpu {
	return &cpu{make(map[string]int)}
}
func (cpu *cpu) run(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	maxRes := 0
	for scanner.Scan() {
		if instr, err := parseLine(scanner.Text()); err == nil {
			res := cpu.eval(instr)
			if res > maxRes {
				maxRes = res
			}
		} else {
			return maxRes, err
		}
	}

	return maxRes, scanner.Err()
}

func parseLine(ln string) (instr, error) {
	tok := strings.Split(ln, " ")

	instrVal, err := strconv.Atoi(tok[instrVal])
	condVal, err := strconv.Atoi(tok[condVal])

	return instr{
		Reg: tok[instrReg],
		Op:  instrOperations[tok[instrOperator]],
		Val: instrVal,
		Cond: cond{
			Reg: tok[condReg],
			Op:  boolOperations[tok[condOperator]],
			Val: condVal,
		},
	}, err
}
func (cpu *cpu) eval(instr instr) int {
	if instr.Cond.Op(cpu.Registers[instr.Cond.Reg], instr.Cond.Val) {
		cpu.Registers[instr.Reg] = instr.Op(cpu.Registers[instr.Reg], instr.Val)
	}
	return cpu.Registers[instr.Reg]
}

func main() {
	cpu := newCPU()
	file, err := os.Open("./day8/day8_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	overallMax, err := cpu.run(io.Reader(file))
	max := 0
	for _, v := range cpu.Registers {
		if v > max {
			max = v
		}
	}

	fmt.Println("-- Day 8 --")
	fmt.Printf("highest_register_value_after=%d, highest_register_value_during_run=%d\n", max, overallMax)
}
