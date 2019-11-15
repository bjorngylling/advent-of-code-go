package operations

type Instr struct {
	Opcode string
	inputA int
	inputB int
	output int
}

func Instruction(opcode string, a, b, output int) Instr {
	return Instr{Opcode: opcode, inputA: a, inputB: b, output: output}
}

type Registers [6]int

func (r Registers) Eq(o Registers) bool {
	for i := range r {
		if r[i] != o[i] {
			return false
		}
	}
	return true
}

type Op func(Instr, Registers) Registers

var Ops = map[string]Op{"muli": muli, "mulr": mulr, "addi": addi, "addr": addr, "bani": bani, "banr": banr,
	"bori": bori, "borr": borr, "seti": seti, "setr": setr, "gtir": gtir, "gtri": gtri, "gtrr": gtrr,
	"eqir": eqir, "eqri": eqri, "eqrr": eqrr}

func muli(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] * instr.inputB
	return reg
}
func mulr(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] * reg[instr.inputB]
	return reg
}

func addi(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] + instr.inputB
	return reg
}

func addr(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] + reg[instr.inputB]
	return reg
}

func bani(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] & instr.inputB
	return reg
}

func banr(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] & reg[instr.inputB]
	return reg
}

func bori(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] | instr.inputB
	return reg
}

func borr(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA] | reg[instr.inputB]
	return reg
}

func seti(instr Instr, reg Registers) Registers {
	reg[instr.output] = instr.inputA
	return reg
}

func setr(instr Instr, reg Registers) Registers {
	reg[instr.output] = reg[instr.inputA]
	return reg
}

func gtir(instr Instr, reg Registers) Registers {
	if instr.inputA > reg[instr.inputB] {
		reg[instr.output] = 1
	} else {
		reg[instr.output] = 0
	}
	return reg
}

func gtri(instr Instr, reg Registers) Registers {
	if reg[instr.inputA] > instr.inputB {
		reg[instr.output] = 1
	} else {
		reg[instr.output] = 0
	}
	return reg
}

func gtrr(instr Instr, reg Registers) Registers {
	if reg[instr.inputA] > reg[instr.inputB] {
		reg[instr.output] = 1
	} else {
		reg[instr.output] = 0
	}
	return reg
}

func eqir(instr Instr, reg Registers) Registers {
	if instr.inputA == reg[instr.inputB] {
		reg[instr.output] = 1
	} else {
		reg[instr.output] = 0
	}
	return reg
}

func eqri(instr Instr, reg Registers) Registers {
	if reg[instr.inputA] == instr.inputB {
		reg[instr.output] = 1
	} else {
		reg[instr.output] = 0
	}
	return reg
}

func eqrr(instr Instr, reg Registers) Registers {
	if reg[instr.inputA] == reg[instr.inputB] {
		reg[instr.output] = 1
	} else {
		reg[instr.output] = 0
	}
	return reg
}
