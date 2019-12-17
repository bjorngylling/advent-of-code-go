package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"strconv"
	"time"
)

type tile = int

const (
	wall tile = iota
	floor
	oxygen
	robot
)

type direction = int

const (
	_ direction = iota
	north
	south
	west
	east
)

var vel = map[direction]image.Point{
	north: image.Pt(0, 1),
	south: image.Pt(0, -1),
	west:  image.Pt(-1, 0),
	east:  image.Pt(1, 0),
}

func generateFrame(world map[image.Point]int, entities map[image.Point]int, palette color.Palette) *image.Paletted {
	r := image.Rect(0, 0, 41, 41)
	img := image.NewPaletted(r, palette)
	for pos, t := range world {
		img.SetColorIndex(pos.X+21, pos.Y+19, uint8(t+1))
	}
	for pos, t := range entities {
		img.SetColorIndex(pos.X+21, pos.Y+19, uint8(t+1))
	}
	return img
}

func writeGif(frames []*image.Paletted) {
	delays := make([]int, len(frames))
	anim := gif.GIF{Delay: delays, Image: frames}
	f, err := os.Create("vis_15.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gif.EncodeAll(f, &anim)
}

func turnLeft(d direction) direction {
	switch d {
	case north:
		return west
	case south:
		return east
	case west:
		return south
	case east:
		return north
	}
	return -1
}

func turnRight(d direction) direction {
	switch d {
	case north:
		return east
	case south:
		return west
	case west:
		return north
	case east:
		return south
	}
	return -1
}

func solve(input string) (string, string) {
	in := make(chan direction)
	out := make(chan tile)
	c, err := intcode.Init(input, in, out)
	if err != nil {
		panic(err)
	}

	go func() {
		c.Run()
		close(out)
	}()

	bot := image.Pt(0, 0)
	oxy := image.Pt(0, 0)
	dir := north
	in <- dir
	world := map[image.Point]tile{image.Pt(0, 0): floor}
	looped := false
	for t := range out {
		switch t {
		case wall:
			p := bot.Add(vel[dir])
			world[p] = t
			if !looped {
				dir = turnLeft(dir)
			} else {
				dir = turnRight(dir)
			}
		case floor:
			world[bot] = t
			bot = bot.Add(vel[dir])
			if looped {
				dir = turnLeft(dir)
			} else {
				dir = turnRight(dir)
			}
		case oxygen:
			bot = bot.Add(vel[dir])
			world[bot] = t
			oxy = bot
			if looped {
				dir = turnLeft(dir)
			} else {
				dir = turnRight(dir)
			}
		}
		select {
		case in <- dir:
		}
		if bot.X == 0 && bot.Y == 0 {
			if oxy.X != 0 && oxy.Y != 0 {
				c.SigInt <- true
			} else {
				looped = true
			}
		}
	}
	if true {
		palette := []color.Color{
			color.Black,
			color.White,
			color.Gray16{Y: 0x3333},
			color.RGBA{R: 255, G: 0, B: 0, A: 255},
			color.RGBA{R: 0, G: 0, B: 255, A: 255},
		}
		var frames []*image.Paletted
		frames = append(frames, generateFrame(world, map[image.Point]int{bot: robot, oxy: oxygen}, palette))
		writeGif(frames)
	}

	part1Cost, _ := dijkstra(world, image.Pt(0, 0))
	part2Cost, _ := dijkstra(world, oxy)
	max := 0
	for _, c := range part2Cost {
		if c > max {
			max = c
		}
	}
	return strconv.Itoa(part1Cost[oxy]), strconv.Itoa(max)
}

func neighbours(p image.Point) [4]image.Point {
	return [4]image.Point{p.Sub(image.Pt(0, 1)), p.Sub(image.Pt(1, 0)),
		p.Add(image.Pt(1, 0)), p.Add(image.Pt(0, 1))}
}

func dijkstra(world map[image.Point]tile, source image.Point) (map[image.Point]int, map[image.Point]image.Point) {
	q := []image.Point{source}                // List containing all unvisited positions
	dist := make(map[image.Point]int)         // Position -> cost to move there from source
	prev := make(map[image.Point]image.Point) // Position -> previous position

	// Add all unblocked positions to q and set the distance there to "infinity"
	for p, t := range world {
		if t == floor || t == oxygen {
			q = append(q, p)
			dist[p] = math.MaxInt32
		}
	}
	dist[source] = 0
	for len(q) > 0 {
		// Find the unvisited position with the lowest distance from source
		iu := -1
		for i, pos := range q {
			if iu == -1 || dist[pos] < dist[q[iu]] {
				iu = i
			}
		}
		u := q[iu]
		q = append(q[:iu], q[iu+1:]...)

		// Check all neighbours of u if there is a shorter path there
		for _, v := range neighbours(u) {
			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	// Remove unreachable positions
	for k, v := range dist {
		if v == math.MaxInt32 {
			delete(dist, k)
			delete(prev, k)
		}
	}
	return dist, prev
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
