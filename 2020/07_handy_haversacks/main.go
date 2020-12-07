package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type pair struct {
	name  string
	count int
}
type node struct {
	name     string
	parent   []string
	children []pair
}

func solve(input string) (string, string) {
	graph := map[string]*node{}
	for _, ln := range strings.Split(input, "\n") {
		w := strings.Fields(ln)
		n := &node{name: strings.Join(w[0:2], " ")}
		for i := 0; i < len(w[4:])/4; i++ {
			idx := i*4 + 4
			c, err := strconv.Atoi(w[idx])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			n.children = append(n.children, pair{name: strings.Join(w[idx+1:idx+3], " "), count: c})
		}
		graph[n.name] = n
	}
	for k, n := range graph {
		for _, c := range n.children {
			graph[c.name].parent = append(graph[c.name].parent, k)
		}
	}

	set := map[string]struct{}{}
	var v func(n *node)
	v = func(n *node) {
		set[n.name] = struct{}{}
		if len(n.parent) == 0 {
			return
		}
		for _, p := range n.parent {
			v(graph[p])
		}
	}
	v(graph["shiny gold"])
	solution1 := len(set) - 1

	var u func(n *node) int
	u = func(n *node) int {
		set[n.name] = struct{}{}
		if len(n.children) == 0 {
			return 1
		}
		sum := 1
		for _, c := range n.children {
			sum += c.count * u(graph[c.name])
		}
		return sum
	}
	solution2 := u(graph["shiny gold"]) - 1

	return strconv.Itoa(solution1), strconv.Itoa(solution2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
