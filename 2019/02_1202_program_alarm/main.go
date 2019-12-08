package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"strconv"
	"time"
)

func solve(input string) (string, string) {
	c1, err := intcode.Init(input, nil, nil)
	if err != nil {
		panic(err)
	}
	c1.Mem[1] = 12
	c1.Mem[2] = 2
	c1.Run()

	part2 := ""
outer:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			c2, err := intcode.Init(input, nil, nil)
			if err != nil {
				panic(err)
			}
			c2.Mem[1] = noun
			c2.Mem[2] = verb
			c2.Run()
			if c2.Mem[0] == 19690720 {
				part2 = fmt.Sprintf("%d%d", c2.Mem[1], c2.Mem[2])
				break outer
			}
		}
	}

	return strconv.Itoa(c1.Mem[0]), part2
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
