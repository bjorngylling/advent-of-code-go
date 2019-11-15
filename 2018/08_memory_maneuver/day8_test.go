package main

import (
	"reflect"
	"testing"
)

var data = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func TestParseLine(t *testing.T) {
	result := parseLine(data)
	expected := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}

func TestBuildTree(t *testing.T) {
	fields := parseLine(data)
	c, result := buildTree(fields)
	if c != len(fields) {
		t.Errorf("All fields were not consumed, consumed %d but should have been %d\n", c, len(fields))
	}
	if len(result.Edges) != 2 {
		t.Errorf("Expected root node to have 2 children but was %d %+v\n", len(result.Edges), result.Edges)
	}
	expected := []int{1, 1, 2}
	if !reflect.DeepEqual(result.Data, expected) {
		t.Errorf("Expected root node data to be %+v but was %+v\n", expected, result.Data)
	}
}

func TestSumData(t *testing.T) {
	_, rootNode := buildTree(parseLine(data))
	result := sumData(rootNode)
	expected := 138
	if result != expected {
		t.Errorf("Expected sumData to be %d but was %d\n", expected, result)
	}
}

func TestNodeValue(t *testing.T) {
	_, rootNode := buildTree(parseLine(data))
	result := nodeValue(rootNode)
	expected := 66
	if result != expected {
		t.Errorf("Expected nodeValue to be %d but was %d\n", expected, result)
	}
}
