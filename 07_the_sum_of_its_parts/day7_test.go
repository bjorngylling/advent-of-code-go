package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	data := "Step C must be finished before step A can begin."
	c, a := parseLine(data)
	if c != "C" && a != "A" {
		t.Errorf("Expected c=C and a=A but was c=%s and a=%s", c, a)
	}
}

var data = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

func TestParseInput(t *testing.T) {
	result := parseInput(data)["C"] // We know C is the only root in the graph described in data
	if result.Name != "C" {
		t.Errorf("Expected root node to be C but was %s", result)
	}
	if len(result.Edges) != 2 {
		t.Errorf("Expected root node to have 2 edges but was %d", len(result.Edges))
	}
	if result.Edges[0].Name != "A" && result.Edges[1].Name != "F" {
		t.Errorf("Expected root node to have edges to A and F but was %s and %s",
			result.Edges[0].Name, result.Edges[1].Name)
	}
}

func TestFindRootNodes(t *testing.T) {
	result := findRootNodes(parseInput(data), make(map[*Node]struct{}))
	if result[0].Name != "C" {
		t.Errorf("Expected root node to be C but was %+v", result)
	}

	result = findRootNodes(parseInput(append(data, "Step Q must be finished before step A can begin.")),
		make(map[*Node]struct{}))
	if result[0].Name != "C" && result[1].Name != "Q" {
		t.Errorf("Expected root nodes to be C and Q but was %+v", nodeListToString(result))
	}
}

func TestWorkOrder(t *testing.T) {
	result := nodeListToString(workOrder(parseInput(data)))
	expected := "CABDFE"
	if result != expected {
		t.Errorf("Expected step order to be %q but was %q", expected, result)
	}
	result = nodeListToString(workOrder(
		parseInput(append(data, "Step Q must be finished before step A can begin."))))
	expected = "CFQABDE"
	if result != expected {
		t.Errorf("Expected step order to be %q but was %q", expected, result)
	}
}

func TestEstimateTime(t *testing.T) {
	result := estimateTime(parseInput(data), 2, 0)
	expected := 15
	if result != expected {
		t.Errorf("Expected estimated time to be %d but was %d", expected, result)
	}

	result = estimateTime(parseInput(data), 2, 1)
	expected = 20
	if result != expected {
		t.Errorf("Expected estimated time to be %d but was %d", expected, result)
	}
	result = estimateTime(parseInput(data), 2, 5)
	expected = 38
	if result != expected {
		t.Errorf("Expected estimated time to be %d but was %d", expected, result)
	}
}

func TestPart1(t *testing.T) {
	fileContent, err := ioutil.ReadFile("day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := parseInput(strings.Split(string(fileContent), "\n"))
	expected := "BCADPVTJFZNRWXHEKSQLUYGMIO"
	result := nodeListToString(workOrder(data))
	if result != expected {
		t.Errorf("Part 1 is wrong, expected %q but was %q", expected, result)
	}
}