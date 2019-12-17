package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"math"
	"strconv"
	"strings"
	"time"
)

type node struct {
	name  string
	count int
	req   []*edge
	total int
}

type edge struct {
	target *node
	cost   int
}

func buildTree(input string) map[string]*node {
	nodes := make(map[string]*node)
	for _, ln := range strings.Split(input, "\n") {
		s := strings.Split(ln, " => ")
		var name string
		var count int
		fmt.Sscanf(s[1], "%d %s", &count, &name)
		if n, found := nodes[name]; found {
			n.count = count
		} else {
			nodes[name] = &node{name: name, count: count, req: nil}
		}
		for _, in := range strings.Split(s[0], ", ") {
			var reqName string
			var reqCost int
			fmt.Sscanf(in, "%d %s", &reqCost, &reqName)
			if _, found := nodes[reqName]; !found {
				nodes[reqName] = &node{name: reqName}
			}
			nodes[name].req = append(nodes[name].req, &edge{target: nodes[reqName], cost: reqCost})
		}
	}
	return nodes
}

func revTopologialSort(n *node) []*node {
	var res []*node
	marked := make(map[*node]struct{})
	var visit func(n *node)
	visit = func(n *node) {
		if _, found := marked[n]; found {
			return
		}
		for _, e := range n.req {
			visit(e.target)
		}
		marked[n] = struct{}{}
		res = append([]*node{n}, res...)
	}
	visit(n)
	return res
}

func part1(input string, reqOre int) int {
	nodes := buildTree(input)
	nodes["FUEL"].total = reqOre

	sorted := revTopologialSort(nodes["FUEL"])
	for _, n := range sorted {
		req := util.Max(1, util.Trunc(math.Ceil(float64(n.total)/float64(n.count))))
		for _, e := range n.req {
			e.target.total += req * e.cost
		}
	}

	return nodes["ORE"].total
}

func part2(input string) int {
	res := 0
	limit := 1000000000000
	right := limit
	left := 1
	for left <= right {
		mid := (left + right) / 2
		reqOre := part1(input, mid)
		if reqOre < limit {
			res = util.Max(res, mid)
			left = mid + 1
		} else if reqOre > limit {
			right = mid - 1
		} else {
			return mid
		}
	}
	return res
}

func solve(input string) (string, string) {
	return strconv.Itoa(part1(input, 1)), strconv.Itoa(part2(input))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
