package grid

import (
	"image"
	"reflect"
	"testing"
)

func TestGrid_ViewPort(t *testing.T) {
	grid := New(200, 200)
	grid.Set(100, 3, Clay)
	grid.Set(100, 4, Clay)
	grid.Set(100, 5, Clay)
	grid.Set(101, 5, Clay)
	grid.Set(105, 5, Clay)
	grid.Set(106, 4, WetSand)
	view := grid.ViewPort(0)
	expected := image.Rect(100, 3, 106, 5)
	if !reflect.DeepEqual(view, expected) {
		t.Errorf("Expected view to be %+v but was %+v", expected, view)
	}
	view = grid.ViewPort(2)
	expected = image.Rect(98, 1, 108, 7)
	if !reflect.DeepEqual(view, expected) {
		t.Errorf("Expected view to be %+v but was %+v", expected, view)
	}
	view = grid.ViewPort(5)
	expected = image.Rect(95, 0, 111, 10)
	if !reflect.DeepEqual(view, expected) {
		t.Errorf("Expected view to be %+v but was %+v", expected, view)
	}
}
