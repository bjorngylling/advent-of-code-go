package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func run(program []string) (int, error) {
	var executed []int
	pntr := 0
	acc := 0
	for {
		if pntr >= len(program) {
			break
		}
		s := strings.Fields(program[pntr])
		n, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, i := range executed {
			if pntr == i {
				return acc, errors.New("infinite loop")
			}
		}
		executed = append(executed, pntr)
		switch s[0] {
		case "jmp":
			pntr += n
			break
		case "acc":
			acc += n
			pntr++
			break
		case "nop":
			pntr++
			break
		}
	}
	return acc, nil
}

func solve(input string) (string, string) {
	split := strings.Split(input, "\n")
	solution1, _ := run(split)

	var solution2 int
	for i := range split {
		p := make([]string, len(split))
		copy(p, split)
		if strings.HasPrefix(split[i], "jmp") {
			p[i] = strings.Replace(split[i], "jmp", "nop", 1)
		} else if strings.HasPrefix(split[i], "nop") {
			p[i] = strings.Replace(split[i], "nop", "jmp", 1)
		}
		r, err := run(p)
		if err == nil {
			solution2 = r
			break
		}
	}
	return strconv.Itoa(solution1), strconv.Itoa(solution2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
