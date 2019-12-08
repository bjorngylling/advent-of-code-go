package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"strconv"
	"time"
)

func solve(input string) (string, string) {
	c1In := make(chan int, 10)
	c1In <- 1
	c1Out := make(chan int, 10)
	c1, err := intcode.Init(input, c1In, c1Out)
	if err != nil {
		panic(err)
	}
	c1.Run()
	close(c1Out)
	part1 := 0
	for diagCode := range c1Out {
		part1 = diagCode
	}

	c2In := make(chan int, 10)
	c2In <- 5
	c2Out := make(chan int, 10)
	c2, err := intcode.Init(input, c2In, c2Out)
	if err != nil {
		panic(err)
	}
	c2.Run()

	return strconv.Itoa(part1), strconv.Itoa(<-c2Out)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
