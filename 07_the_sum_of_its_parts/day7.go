package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Node struct {
	Name  string
	Edges []Node
}

func NewNode(name string) Node {
	return Node{Name: name}
}

func parseInput(lines []string) Node {
	lookupTbl := make(map[string]Node)
	for ln := range lines {

	}
}

func main() {
	fileContent, err := ioutil.ReadFile("07_the_sum_of_its_parts/day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parseInput(strings.Split(string(fileContent), "\n"))

	fmt.Printf("Day 6 part 1 result: %+v\n", nil)

	fmt.Printf("Day 6 part 2 result: %+v\n", nil)
}
