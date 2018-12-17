package main

import (
	"reflect"
	"testing"
)

func TestScores(t *testing.T) {
	r := scores(5)
	e := []int{0, 1, 2, 4, 5, 1, 5, 8, 9, 1}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected the ten scores after 5 iterations to be %+v but was %+v", e, r)
	}
	r = scores(18)
	e = []int{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected the ten scores after 5 iterations to be %+v but was %+v", e, r)
	}
	r = scores(2018)
	e = []int{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected the ten scores after 5 iterations to be %+v but was %+v", e, r)
	}
	r = scores(236021)
	e = []int{6, 2, 9, 7, 3, 1, 0, 8, 6, 2}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected the ten scores after 5 iterations to be %+v but was %+v", e, r)
	}
}

func TestRounds(t *testing.T) {
	r := rounds([]int{0, 1, 2, 4, 5})
	e := 5
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected %+v but was %+v", e, r)
	}
	r = rounds([]int{5, 1, 5, 8, 9})
	e = 9
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected %+v but was %+v", e, r)
	}
	r = rounds([]int{9, 2, 5, 1, 0})
	e = 18
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected %+v but was %+v", e, r)
	}
	r = rounds([]int{5, 9, 4, 1, 4})
	e = 2018
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected %+v but was %+v", e, r)
	}
	r = rounds([]int{2, 3, 6, 0, 2, 1})
	e = 20221334
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected %+v but was %+v", e, r)
	}
}
