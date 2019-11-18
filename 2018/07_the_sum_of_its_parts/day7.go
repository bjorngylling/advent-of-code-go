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

func findRootNodes(nodeTbl map[string]*Node, ignore map[*Node]struct{}) []*Node {
	rootNodes := make(map[string]struct{})
	for k, n := range nodeTbl {
		if _, ok := ignore[n]; !ok {
			rootNodes[k] = struct{}{}
		}
	}
	for _, n := range nodeTbl {
		if _, ok := ignore[n]; !ok {
			for _, e := range n.Edges {
				delete(rootNodes, e.Name)
			}
		}
	}
	var r []*Node
	for k := range rootNodes {
		r = append(r, nodeTbl[k])
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
	handled := make(map[*Node]struct{})
	rootNodes := findRootNodes(nodeTbl, handled)
	for len(rootNodes) > 0 {
		k := rootNodes[0]
		rootNodes = rootNodes[1:]
		handled[k] = struct{}{}
		l = append(l, k)
		rootNodes = findRootNodes(nodeTbl, handled)
	}
	return
}

type Task struct {
	step     *Node
	timeLeft int
}

func (t *Task) process() {
	t.timeLeft -= 1
}

func estimateTime(nodeTbl map[string]*Node, workerCount int, baseTime int) (t int) {
	completed := make(map[*Node]struct{})
	assigned := make(map[*Node]struct{})
	workers := make([]*Task, workerCount)
	steps := findRootNodes(nodeTbl, completed)
	// while there are steps to complete
	for len(completed) < len(nodeTbl) {
		// assign steps to workers
		for i, w := range workers {
			if w == nil { // only idle workers
				for _, s := range steps { // find unassigned step and assign it
					if _, ok := assigned[s]; !ok {
						workers[i] = &Task{s, baseTime + int(s.Name[0]) - 64}
						assigned[s] = struct{}{}
						break // move on to the next worker
					}
				}
			}
		}
		// process all workers
		for _, w := range workers {
			if w != nil {
				w.process()
			}
		}
		// free up completed workers
		for i, w := range workers {
			if w != nil && w.timeLeft <= 0 {
				workers[i] = nil
				completed[w.step] = struct{}{}
			}
		}
		steps = findRootNodes(nodeTbl, completed)
		t += 1
	}
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/07_the_sum_of_its_parts/day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 7 part 1 result: %+v\n", nodeListToString(workOrder(
		parseInput(strings.Split(string(fileContent), "\n")))))

	fmt.Printf("Day 7 part 2 result: %+v\n", estimateTime(
		parseInput(strings.Split(string(fileContent), "\n")), 5, 60))
}
