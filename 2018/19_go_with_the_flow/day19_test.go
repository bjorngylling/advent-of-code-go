package main

import (
	. "github.com/bjorngylling/advent-of-code-2018/19_go_with_the_flow/operations"
	"testing"
)

var input = `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`

func TestParseProgram(t *testing.T) {
	ip, program := parseInput(input)

	want := 0
	if ip != want {
		t.Errorf("ip == %d, want %d", ip, want)
	}
	want = 7
	got := len(program)
	if got != want {
		t.Errorf("ip == %d, want %d", got, want)
	}
}

func TestRunProgram(t *testing.T) {
	ip, program := parseInput(input)
	want := Registers{6, 5, 6, 0, 0, 9}
	got := runProgram(ip, program)
	if !got.Eq(want) {
		t.Errorf("Registers == %d, want %d", got, want)
	}
}
