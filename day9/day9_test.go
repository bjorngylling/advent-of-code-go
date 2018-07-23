package main

import (
"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestBasicGroups(t *testing.T) {
	n, _ := process("{}")
	assertEqual(t, 1, n)
	n, _ = process("{{}}")
	assertEqual(t, 3, n)
	n, _ = process("{{{}}}")
	assertEqual(t, 6, n)
	n, _ = process("{{},{}}")
	assertEqual(t, 5, n)
	n, _ = process("{{{},{},{{}}}}")
	assertEqual(t, 16, n)
}
func TestGroupsWithGarbage(t *testing.T) {
	n, _ := process("{<a>,<a>,<a>,<a>}")
	assertEqual(t, 1, n)
	n, _ = process("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	assertEqual(t, 9, n)
	n, _ = process("{{<ab>},{<{}ab>},{<}ab>},{<ab>}}")
	assertEqual(t, 9, n)
}
func TestGroupsWithEscapeChar(t *testing.T) {
	n, _ := process("{{<!!>},{<!!>},{<!!>},{<!!>}}")
	assertEqual(t, 9, n)
	n, _ = process("{{}}!{}")
	assertEqual(t, 3, n)
	n, _ = process("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	assertEqual(t, 3, n)
}
func TestGarbageCount(t *testing.T) {
	_, n := process("{{}}!{}")
	assertEqual(t, 0, n)
	_, n = process("<{!>}>")
	assertEqual(t, 2, n)
	_, n = process("{{<!!>},{<!!>},{<!!>},{<!!>}}")
	assertEqual(t, 0, n)
	_, n = process("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	assertEqual(t, 17, n)
}