package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type Node struct {
	Name  string
	Edges []*Node
}

func NewNode(name string) *Node {
	return &Node{Name: name, Edges: make([]*Node, 0)}
}

func (n *Node) AddEdge(other *Node) {
	n.Edges = append(n.Edges, other)
}

func parseLine(ln string) (string, string) {
	f := strings.Fields(ln)
	return f[1], f[7]
}

func parseInput(lines []string) map[string]*Node {
	lookupTbl := make(map[string]*Node)
	for _, ln := range lines {
		srcName, tgtName := parseLine(ln)
		tgt, ok := lookupTbl[tgtName]
		if !ok {
			tgt = NewNode(tgtName)
			lookupTbl[tgtName] = tgt
		}
		src, ok := lookupTbl[srcName]
		if !ok {
			src = NewNode(srcName)
			lookupTbl[srcName] = src
		}
		src.AddEdge(tgt)
	}
	return lookupTbl
}

func findRootNodes(nodeTbl map[string]*Node) []*Node {
	rootNodes := make(map[string]*Node)
	for k, n := range nodeTbl {
		rootNodes[k] = n
	}
	for _, n := range nodeTbl {
		for _, e := range n.Edges {
			delete(rootNodes, e.Name)
		}
	}
	var r []*Node
	for _, n := range rootNodes {
		r = append(r, n)
	}
	sortNodeList(r)
	return r
}

func nodeListToString(l []*Node) (s string) {
	for _, n := range l {
		s += n.Name
	}
	return
}

func sortNodeList(l []*Node) {
	sort.Slice(l, func(i, j int) bool {
		return strings.Compare(l[i].Name, l[j].Name) == -1
	})
}

func workOrder(nodeTbl map[string]*Node) (l []*Node) {
	rootNodes := findRootNodes(nodeTbl)
	for len(rootNodes) > 0 {
		k := rootNodes[0]
		rootNodes = rootNodes[1:]
		delete(nodeTbl, k.Name)
		l = append(l, k)
		rootNodes = findRootNodes(nodeTbl)
	}
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("07_the_sum_of_its_parts/day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := parseInput(strings.Split(string(fileContent), "\n"))

	fmt.Printf("Day 7 part 1 result: %+v\n", nodeListToString(workOrder(data)))

	fmt.Printf("Day 7 part 2 result: %+v\n", nil)
}
