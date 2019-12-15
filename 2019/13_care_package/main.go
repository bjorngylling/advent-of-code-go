package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"image"
	"image/color"
	"image/gif"
	"os"
	"strconv"
	"time"
)

const (
	empty = iota
	wall
	block
	paddle
	ball
)

func part1(input string) int {
	in := make(chan int)
	out := make(chan int)
	c, err := intcode.Init(input, in, out)
	if err != nil {
		panic(err)
	}

	go func() {
		c.Run()
		close(out)
	}()

	world := make(map[image.Point]int)
	for x := range out {
		y := <-out
		t := <-out

		world[image.Pt(x, y)] = t
	}

	return countBlocks(world)
}

func countBlocks(world map[image.Point]int) int {
	blockCount := 0
	for _, e := range world {
		if e == block {
			blockCount++
		}
	}
	return blockCount
}

func generateFrame(world map[image.Point]int) *image.Paletted {
	palette := []color.Color{
		color.Black,
		color.Gray16{Y: 0x7777},
		color.White,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		color.RGBA{R: 0, G: 0, B: 255, A: 255}}
	rect := image.Rect(0, 0, 40, 25)
	img := image.NewPaletted(rect, palette)
	for pos, t := range world {
		img.SetColorIndex(pos.X, pos.Y, uint8(t))
	}
	return img
}

func writeGif(frames []*image.Paletted) {
	delays := make([]int, len(frames))
	anim := gif.GIF{Delay: delays, Image: frames}
	f, err := os.Create("vis_13_2.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gif.EncodeAll(f, &anim)
}

func part2(input string) int {
	in := make(chan int, 1)
	in <- 0
	out := make(chan int)
	c, err := intcode.Init(input, in, out)
	if err != nil {
		panic(err)
	}
	c.Mem[0] = 2

	go func() {
		c.Run()
		close(out)
	}()

	world := make(map[image.Point]int)
	p := image.Pt(0, 0)
	score := 0

	var frames []*image.Paletted
	gameStarted := false
	for x := range out {
		y := <-out
		t := <-out

		switch {
		case x == -1 && y == 0:
			score = t
		default:
			world[image.Pt(x, y)] = t
		}

		// The game starts when the ball begins to move
		gameStarted = gameStarted || (t == ball && x == 19)

		if t == paddle {
			p.X, p.Y = x, y
		}

		// Only generate input when the ball is updated
		if gameStarted && t == ball {
			input := 0
			switch {
			case p.X < x:
				input++
			case p.X > x:
				input--
			}
			in <- input
		}

		// Generate an image of the current state for the GIF
		if t == ball {
			frames = append(frames, generateFrame(world))
		}
	}
	writeGif(frames)

	return score
}

func solve(input string) (string, string) {
	return strconv.Itoa(part1(input)), strconv.Itoa(part2(input))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
