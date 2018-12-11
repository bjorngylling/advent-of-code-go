package main

import (
	"testing"
)

func TestPowerLevel(t *testing.T) {
	result := powerLevel(3, 5, 8)
	expected := 4
	if result != expected {
		t.Errorf("Expected power level for [x=3, y=5 gridSerialNo=8] to be %d, but was %d", expected, result)
	}
	result = powerLevel(122, 79, 57)
	expected = -5
	if result != expected {
		t.Errorf("Expected power level for [x=122, y=79 gridSerialNo=57] to be %d, but was %d", expected, result)
	}
	result = powerLevel(217, 196, 39)
	expected = 0
	if result != expected {
		t.Errorf("Expected power level for [x=217, y=196 gridSerialNo=0] to be %d, but was %d", expected, result)
	}
	result = powerLevel(101, 153, 71)
	expected = 4
	if result != expected {
		t.Errorf("Expected power level for [x=101, y=153 gridSerialNo=71] to be %d, but was %d", expected, result)
	}
}

func TestGroupPowerLevel(t *testing.T) {
	result := groupPowerLevel(33, 45, 3, 18)
	expected := 29
	if result != expected {
		t.Errorf("Expected group power level for [x=33, y=45 gridSerialNo=18] to be %d, but was %d", expected, result)
	}
	result = groupPowerLevel(21, 61, 3, 42)
	expected = 30
	if result != expected {
		t.Errorf("Expected group power level for [x=21, y=61 gridSerialNo=42] to be %d, but was %d", expected, result)
	}
}

func TestFindHighestGroupPower(t *testing.T) {
	x, y := findHighestGroupPower(18)
	eX, eY := 33, 45
	if x != eX && y != eY {
		t.Errorf("Expected [x=%d y=%d] but was [x=%d y=%d]", x, y, eX, eY)
	}
}

func TestFindMaxGroupPower(t *testing.T) {
	x, y, s := findMaxGroupPower(18)
	eX, eY, eS := 90, 269, 16
	if x != eX && y != eY && s != eS {
		t.Errorf("Expected [x=%d y=%d s=%d] but was [x=%d y=%d, s%d]", x, y, s, eX, eY, eS)
	}
}
