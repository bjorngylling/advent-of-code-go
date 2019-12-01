package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	part1 := 0
	part2 := 0
	for _, mass := range strings.Split(input, "\n") {
		m, e := strconv.Atoi(mass)
		if e != nil {
			panic(e)
		}
		part1 += m/3 - 2
		part2 += calculateFuelForMass(m)
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func calculateFuelForMass(m int) int {
	fuel := m/3 - 2
	if fuel > 0 {
		fuel += calculateFuelForMass(fuel)
	} else if fuel < 0 {
		fuel = 0
	}
	return fuel
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
