package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"strconv"
	"time"
)

func solve(input string) (string, string) {
	in := make(chan int, 10)
	in <- 1
	out := make(chan int, 100)
	computer, err := intcode.Init(input, in, out)
	if err != nil {
		panic(err)
	}
	computer.Run()

	c2In := make(chan int, 10)
	c2In <- 2
	c2Out := make(chan int, 100)
	c2, err := intcode.Init(input, c2In, c2Out)
	if err != nil {
		panic(err)
	}
	c2.Run()

	return strconv.Itoa(<-out), strconv.Itoa(<-c2Out)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
