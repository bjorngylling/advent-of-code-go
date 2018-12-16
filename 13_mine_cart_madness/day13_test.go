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
	eCarts := []Cart{{Pos: Pos(2, 0), Speed: East}, {Pos: Pos(9, 3), Speed: South}}

	if !reflect.DeepEqual(grid, eGrid) {
		t.Errorf("Expected the grid to be %+v but was %+v", eGrid, grid)
	}
	if !reflect.DeepEqual(carts, eCarts) {
		t.Errorf("Expected the mine carts be %+v but was %+v", eCarts, carts)
	}
}

func TestStep(t *testing.T) {
	grid, carts := parseInput(input)
	step(grid, carts)

	eCarts := []Cart{{Pos: Pos(3, 0), Speed: East}, {Pos: Pos(9, 4), Speed: East, TurnCount: 1}}
	if !reflect.DeepEqual(carts, eCarts) {
		t.Errorf("Expected the mine carts be %+v but was %+v", eCarts, carts)
	}
}

func TestFirstCollision(t *testing.T) {
	coll := firstCollision(parseInput(input))
	ePos := Pos(7, 3)

	if !reflect.DeepEqual(ePos, coll) {
		t.Errorf("Expected the first collision to happen at %+v but was %+v", ePos, coll)
	}
}

func TestLastCart(t *testing.T) {
	input := `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`
	cart := lastCart(parseInput(input))
	ePos := Pos(6, 4)

	if !reflect.DeepEqual(ePos, cart.Pos) {
		t.Errorf("Expected the last cart to be at %+v but was %+v", ePos, cart.Pos)
	}
}
