package main

import (
	"fmt"
	. "github.com/bjorngylling/advent-of-code/2018/19_go_with_the_flow/operations"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) (int, []Instr) {
	x := strings.Split(input, "\n")
	ip := mustAtoi(strings.Fields(x[0])[1])
	var program []Instr
	for _, ln := range x[1:] {
		program = append(program, parseInstr(ln))
	}
	return ip, program
}

func parseInstr(s string) Instr {
	f := strings.Fields(s)
	return Instruction(f[0], mustAtoi(f[1]), mustAtoi(f[2]), mustAtoi(f[3]))
}

func mustAtoi(s string) int {
	if n, err := strconv.Atoi(s); err != nil {
		log.Fatalf("Failed to convert %q to int", s)
	} else {
		return n
	}
	return 0
}

func runProgram(ipIdx int, program []Instr, reg Registers) Registers {
	for ip := reg[ipIdx]; ip >= 0 && ip < len(program); {
		reg[ipIdx] = ip
		reg = Ops[program[ip].Opcode](program[ip], reg)
		ip = reg[ipIdx] + 1
	}
	return reg
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/19_go_with_the_flow/day19_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ip, program := parseInput(string(fileContent))

	fmt.Printf("Day 19 part 1 result: %+v\n", runProgram(ip, program, Registers{})[0])
}
