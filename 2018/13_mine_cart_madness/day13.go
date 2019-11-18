package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type Position = image.Point
type Speed = image.Point

var North = Speed{Y: -1}
var East = Speed{X: 1}
var South = Speed{Y: 1}
var West = Speed{X: -1}

var Empty = Pos(-1, -1)

var RightHand = map[Speed]Speed{
	North: East,
	East:  South,
	South: West,
	West:  North,
}
var LeftHand = map[Speed]Speed{
	North: West,
	East:  North,
	South: East,
	West:  South,
}

type Cart struct {
	Pos       Position
	Speed     Speed
	TurnCount int
}

func Pos(x, y int) Position {
	return Position{X: x, Y: y}
}

func parseInput(input string) ([]string, []Cart) {
	lines := strings.Split(input, "\n")
	var carts []Cart
	for y, ln := range lines {
		for x, c := range ln {
			switch c {
			case rune('^'):
				carts = append(carts, Cart{Pos: Pos(x, y), Speed: North})
			case rune('>'):
				carts = append(carts, Cart{Pos: Pos(x, y), Speed: East})
			case rune('v'):
				carts = append(carts, Cart{Pos: Pos(x, y), Speed: South})
			case rune('<'):
				carts = append(carts, Cart{Pos: Pos(x, y), Speed: West})
			}
		}
		lines[y] = strings.Replace(lines[y], "^", "|", -1)
		lines[y] = strings.Replace(lines[y], "v", "|", -1)
		lines[y] = strings.Replace(lines[y], ">", "-", -1)
		lines[y] = strings.Replace(lines[y], "<", "-", -1)
	}
	return lines, carts
}

func step(rail []string, carts []Cart) (collisions []Position) {
	sort.Slice(carts, func(i, j int) bool {
		return carts[i].Pos.Y < carts[j].Pos.Y ||
			(carts[i].Pos.Y == carts[j].Pos.Y && carts[i].Pos.X < carts[j].Pos.X)
	})
OUTER:
	for i := range carts {
		// Skip any carts involved in a collision
		for _, c := range collisions {
			if carts[i].Pos.Eq(c) {
				continue OUTER
			}
		}

		// Calculate the carts new position
		pos := carts[i].Pos.Add(carts[i].Speed)

		// Check for collisions
		for _, other := range carts {
			if pos.Eq(other.Pos) {
				collisions = append(collisions, pos)
			}
		}
		carts[i].Pos = pos

		// Calculate new speed
		switch rail[carts[i].Pos.Y][carts[i].Pos.X] {
		case '/':
			if carts[i].Speed.Eq(North) || carts[i].Speed.Eq(South) {
				carts[i].Speed = RightHand[carts[i].Speed]
			} else {
				carts[i].Speed = LeftHand[carts[i].Speed]
			}
		case '\\':
			if carts[i].Speed.Eq(North) || carts[i].Speed.Eq(South) {
				carts[i].Speed = LeftHand[carts[i].Speed]
			} else {
				carts[i].Speed = RightHand[carts[i].Speed]
			}
		case '+':
			switch carts[i].TurnCount % 3 {
			case 0:
				carts[i].Speed = LeftHand[carts[i].Speed]
			case 2:
				carts[i].Speed = RightHand[carts[i].Speed]
			default:
			}
			carts[i].TurnCount += 1
		}
	}
	return collisions
}

func firstCollision(rail []string, carts []Cart) Position {
	coll := step(rail, carts)
	for ; len(coll) == 0; coll = step(rail, carts) {
	}
	return coll[0]
}

func lastCart(rail []string, carts []Cart) Cart {
	collisions := step(rail, carts)
	for true {
		// Remove all colliding carts
		for _, c := range collisions {
			for i := len(carts) - 1; i >= 0; i-- {
				if carts[i].Pos.Eq(c) {
					carts = append(carts[:i], carts[i+1:]...)
				}
			}
		}
		if len(carts) <= 1 {
			break
		}
		collisions = step(rail, carts)
	}
	return carts[0]
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/13_mine_cart_madness/day13_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 13 part 1 result: %+v\n", firstCollision(parseInput(string(fileContent))))

	fmt.Printf("Day 13 part 2 result: %+v\n", lastCart(parseInput(string(fileContent))).Pos)
}
