package main

import (
	"image"
	"reflect"
	"testing"
)

func TestCreateWorld(t *testing.T) {
	data := []string{"#1 @ 1,3: 4x4"}
	world := createWorld(data)
	expectedWorld := make(World)
	expectedWorld[1] = image.Rect(1, 3, 5, 7)
	if !reflect.DeepEqual(world, expectedWorld) {
		t.Errorf("Expected %+v but was %+v", expectedWorld, world)
	}
}

func TestContestedPositions(t *testing.T) {
	data := []string{"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
		"#4 @ 6,5: 2x2",
		"#5 @ 1,1: 2x2"}
	world := createWorld(data)
	n := contestedPoints(world)

	expected := make(map[image.Point]struct{})
	expected[image.Pt(3, 3)] = struct{}{}
	expected[image.Pt(3, 4)] = struct{}{}
	expected[image.Pt(4, 3)] = struct{}{}
	expected[image.Pt(4, 4)] = struct{}{}
	expected[image.Pt(6, 5)] = struct{}{}
	expected[image.Pt(6, 6)] = struct{}{}
	if !reflect.DeepEqual(n, expected) {
		t.Errorf("Expected 4 overlapping squares, %+v but was %+v", expected, n)
	}
}

func TestUncontestedClaim(t *testing.T) {
	data := []string{"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
		"#4 @ 6,5: 2x2",
		"#5 @ 1,1: 2x2"}
	world := createWorld(data)
	n := uncontestedClaims(world)

	if len(n) != 1 && n[0] != 5 {
		t.Errorf("Expected uncontested claim to be [5] but was %v", n)
	}
}
