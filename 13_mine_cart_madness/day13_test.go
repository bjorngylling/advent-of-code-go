package main

import (
	"reflect"
	"testing"
)

var input = `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

func TestParseInput(t *testing.T) {
	grid, carts := parseInput(input)
	eGrid := []string{
		`/---\        `,
		`|   |  /----\`,
		`| /-+--+-\  |`,
		`| | |  | |  |`,
		`\-+-/  \-+--/`,
		`  \------/   `}
	eCarts := []Cart{{Pos(2, 0), East}, {Pos(9, 3), South}}

	if !reflect.DeepEqual(grid, eGrid) {
		t.Errorf("Expected the grid to be %+v but was %+v", eGrid, grid)
	}
	if !reflect.DeepEqual(carts, eCarts) {
		t.Errorf("Expected the mine carts be %+v but was %+v", eCarts, carts)
	}
}
