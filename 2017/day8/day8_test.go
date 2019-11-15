package main

import (
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	cpu := newCPU()
	if _, err := cpu.run(strings.NewReader(sampleInput)); err != nil {
		t.Error(err)
	}

	if cpu.Registers["a"] != 1 {
		t.Errorf("Expected a to be 1, but was %d instead", cpu.Registers["a"])
	}
	if cpu.Registers["b"] != 0 {
		t.Errorf("Expected b to be 0, but was %d instead", cpu.Registers["b"])
	}
	if cpu.Registers["c"] != -10 {
		t.Errorf("Expected c to be -10, but was %d instead", cpu.Registers["c"])
	}
}

const sampleInput = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
