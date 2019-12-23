package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"image"
	"strconv"
	"time"
)

func beamActiveAt(input []int, x int, y int) bool {
	in := make(chan int)
	out := make(chan int)
	mem := make([]int, len(input))
	copy(mem, input)
	c := intcode.Computer{Mem: mem, In: in, Out: out, SigInt: make(chan bool, 1)}

	go func() {
		c.Run()
		close(out)
	}()

	in <- x
	in <- y
	res := <-out == 1
	c.SigInt <- true
	return res
}

func solve(input string) (string, string) {
	c, err := intcode.Init(input, nil, nil)
	if err != nil {
		panic(err)
	}
	world := map[image.Point]bool{}
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			world[image.Pt(x, y)] = beamActiveAt(c.Mem, x, y)
		}
	}

	total := 0
	for _, b := range world {
		if b {
			total++
		}
	}

	part2 := 0
outer:
	for y := 0; ; y++ {
		for x := 0; ; x++ {
			if beamActiveAt(c.Mem, x, y+99) {
				if beamActiveAt(c.Mem, x+99, y) {
					part2 = x*10000 + y
					break outer
				}
				break
			}
		}
	}

	return strconv.Itoa(total), strconv.Itoa(part2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
