package main

import (
	"reflect"
	"testing"
)

func TestRounds(t *testing.T) {
	r := 0
	e := 5
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected %+v but was %+v", e, r)
	}
}
