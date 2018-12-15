package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"strings"
)

type Position image.Point
type Speed image.Point

var North = Speed{Y: -1}
var East = Speed{X: 1}
var South = Speed{Y: 1}
var West = Speed{X: -1}

type Cart struct {
	pos   Position
	speed Speed
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
				carts = append(carts, Cart{Pos(x, y), North})
			case rune('>'):
				carts = append(carts, Cart{Pos(x, y), East})
			case rune('v'):
				carts = append(carts, Cart{Pos(x, y), South})
			case rune('<'):
				carts = append(carts, Cart{Pos(x, y), West})
			}
		}
		lines[y] = strings.Replace(lines[y], "^", "|", -1)
		lines[y] = strings.Replace(lines[y], "v", "|", -1)
		lines[y] = strings.Replace(lines[y], ">", "-", -1)
		lines[y] = strings.Replace(lines[y], "<", "-", -1)
	}
	return lines, carts
}

func main() {
	fileContent, err := ioutil.ReadFile("13_mine_cart_madness/day13_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 13 part 1 result: %+v\n", fileContent)

	fmt.Printf("Day 13 part 2 result: %+v\n", nil)
}
