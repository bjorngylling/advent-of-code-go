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
}
