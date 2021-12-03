package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type cmd struct {
	dir  string
	dist int
}

func part1(cmds []cmd) int {
	pos, depth := 0, 0
	for _, c := range cmds {
		switch c.dir {
		case "forward":
			pos += c.dist
		case "up":
			depth -= c.dist
		case "down":
			depth += c.dist
		}
	}
	return pos * depth
}

func part2(cmds []cmd) int {
	pos, depth, aim := 0, 0, 0
	for _, c := range cmds {
		switch c.dir {
		case "forward":
			pos += c.dist
			depth += c.dist * aim
		case "up":
			aim -= c.dist
		case "down":
			aim += c.dist
		}
	}
	return pos * depth
}

func solve(input string) (string, string) {
	lns := strings.Split(input, "\n")
	var cmds []cmd
	for _, ln := range lns {
		dir := ""
		dist := 0
		_, err := fmt.Sscanf(ln, "%s %d", &dir, &dist)
		if err != nil {
			panic(err)
		}
		cmds = append(cmds, cmd{dir, dist})
	}

	return strconv.Itoa(part1(cmds)), strconv.Itoa(part2(cmds))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
