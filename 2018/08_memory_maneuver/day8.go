package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	FIELD_CHILD_COUNT = iota
	FIELD_DATA
	HEADER_LEN
)

type Node struct {
	Edges []*Node
	Data  []int
}

func NewNode() *Node {
	return &Node{Edges: make([]*Node, 0)}
}

func (n *Node) AddEdge(other *Node) {
	n.Edges = append(n.Edges, other)
}

func parseLine(ln string) (r []int) {
	for _, s := range strings.Fields(ln) {
		n, _ := strconv.Atoi(s)
		r = append(r, n)
	}
	return
}

func buildTree(data []int) (c int, n *Node) {
	if data[FIELD_CHILD_COUNT] == 0 {
		c = HEADER_LEN + data[FIELD_DATA]
		n = NewNode()
		n.Data = data[HEADER_LEN : HEADER_LEN+data[FIELD_DATA]]
	} else {
		n = NewNode()
		c = HEADER_LEN
		for i := data[FIELD_CHILD_COUNT]; i > 0; i-- {
			d, child := buildTree(data[c:])
			n.AddEdge(child)
			c += d
		}
		n.Data = data[c : c+data[FIELD_DATA]]
		c += data[FIELD_DATA]
	}
	return
}

func sumData(node *Node) (sum int) {
	if len(node.Edges) == 0 {
		for _, d := range node.Data {
			sum += d
		}
	} else {
		for _, n := range node.Edges {
			sum += sumData(n)
		}
		for _, d := range node.Data {
			sum += d
		}
	}
	return
}

func nodeValue(node *Node) (sum int) {
	if len(node.Edges) == 0 {
		for _, d := range node.Data {
			sum += d
		}
	} else {
		for _, d := range node.Data {
			if d <= len(node.Edges) {
				sum += nodeValue(node.Edges[d-1])
			}
		}
	}
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/08_memory_maneuver/day8_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, node := buildTree(parseLine(string(fileContent)))
	fmt.Printf("Day 8 part 1 result: %+v\n", sumData(node))

	fmt.Printf("Day 8 part 2 result: %+v\n", nodeValue(node))
}
