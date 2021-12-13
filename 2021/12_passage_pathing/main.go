package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type node struct {
	name    string
	large   bool
	next    []*node
	visited int
}

func addNode(g map[string]*node, n string) *node {
	if _, ok := g[n]; !ok {
		g[n] = &node{name: n, large: unicode.IsUpper([]rune(n)[0])}
	}
	return g[n]
}

func addEdge(f, t *node) {
	f.next = append(f.next, t)
	t.next = append(t.next, f)
}

func createGraph(input string) (*node, *node) {
	g := map[string]*node{}
	for _, ln := range strings.Split(input, "\n") {
		spl := strings.Split(ln, "-")
		addEdge(addNode(g, spl[0]), addNode(g, spl[1]))
	}
	start, end := g["start"], g["end"]
	start.visited = 1
	return start, end
}

func findPaths(src, dest *node, count int, twice bool) int {
	if src == dest {
		return count + 1
	}
	for _, nn := range src.next {
		if nn.large || nn.visited == 0 || (nn.name != "start" && !twice) {
			nn.visited++
			count = findPaths(nn, dest, count, twice || (!nn.large && nn.visited == 2))
			nn.visited--
		}
	}
	return count
}

func solve(input string) (string, string) {
	start, end := createGraph(input)
	p1 := findPaths(start, end, 0, true)
	start, end = createGraph(input)
	p2 := findPaths(start, end, 0, false)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
