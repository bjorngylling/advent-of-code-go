package main

import (
	"fmt"
	"image"
	"reflect"
	"testing"
)

func TestFindBounds(t *testing.T) {
	pts := []string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"}
	result := findBounds(parseInput(pts))
	expected := image.Rect(1, 1, 8, 9)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}

func TestManhattanDistance(t *testing.T) {
	a, b := image.Pt(5, 0), image.Pt(9, 11)
	expected := 15
	result := manhattanDistance(a, b)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
	result = manhattanDistance(b, a)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}

	a, b = image.Pt(0, 5), image.Pt(11, 9)
	expected = 15
	result = manhattanDistance(a, b)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
	result = manhattanDistance(b, a)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}

	a, b = image.Pt(0, 5), image.Pt(8, 5)
	expected = 8
	result = manhattanDistance(a, b)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
	result = manhattanDistance(b, a)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}

	a, b = image.Pt(0, 9), image.Pt(1, 6)
	expected = 4
	result = manhattanDistance(a, b)
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}

func TestFillGrid(t *testing.T) {
	pts := parseInput([]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"})
	result := fillGrid(pts, findBounds(pts))
	expected := [][]int{
		{0, 0, 0, 0, 0, -1, 2, 2, 2},
		{0, 0, 0, 0, 0, -1, 2, 2, 2},
		{0, 0, 0, 3, 3, 4, 2, 2, 2},
		{0, 0, 3, 3, 3, 4, 2, 2, 2},
		{-1, -1, 3, 3, 3, 4, 4, 2, 2},
		{1, 1, -1, 3, 4, 4, 4, 4, 2},
		{1, 1, 1, -1, 4, 4, 4, 4, -1},
		{1, 1, 1, -1, 4, 4, 4, 5, 5},
		{1, 1, 1, -1, 4, 4, 5, 5, 5},
		{1, 1, 1, -1, 5, 5, 5, 5, 5},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected \n%s but was \n%s", fmtGrid(expected), fmtGrid(result))
	}
}

func TestLargestContainedArea(t *testing.T) {
	pts := parseInput([]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"})
	bounds := findBounds(pts)
	grid := fillGrid(pts, bounds)
	result := largestContainedArea(grid, bounds)
	expected := 17
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	pts := parseInput([]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"})
	bounds := findBounds(pts)
	result := part2(pts, bounds, 32)
	expected := 16
	if expected != result {
		t.Errorf("Expected %d but was %d", expected, result)
	}
}

func fmtGrid(grid [][]int) (s string) {
	for _, row := range grid {
		s += fmt.Sprintf("%+v\n", row)
	}
	return s
}
