package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"image"
	"strconv"
	"strings"
	"sync"
	"time"
)

func neighbours(p image.Point) [4]image.Point {
	return [4]image.Point{p.Sub(image.Pt(0, 1)), p.Sub(image.Pt(1, 0)),
		p.Add(image.Pt(1, 0)), p.Add(image.Pt(0, 1))}
}

func part1(input string) int {
	in := make(chan int)
	out := make(chan int)
	computer, err := intcode.Init(input, in, out)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	s := ""
	go func() {
		for r := range out {
			s += string(r)
		}
		wg.Done()
	}()
	go func() {
		computer.Run()
		close(out)
	}()
	wg.Wait()

	world := map[image.Point]struct{}{}
	for y, ln := range strings.Split(s, "\n") {
		for x, t := range ln {
			if t == '#' {
				world[image.Pt(x, y)] = struct{}{}
			}
		}
	}
	var crossings []image.Point
	for p := range world {
		cross := true
		for _, n := range neighbours(p) {
			if _, found := world[n]; !found {
				cross = false
				break
			}
		}
		if cross {
			crossings = append(crossings, p)
		}
	}
	part1 := 0
	for _, crossing := range crossings {
		part1 += crossing.X * crossing.Y
	}
	return part1
}

func part2(input string) int {
	in := make(chan int, 100)
	out := make(chan int)
	computer, err := intcode.Init(input, in, out)
	if err != nil {
		panic(err)
	}
	computer.Mem[0] = 2

	instr := []string{
		"A,B,A,B,C,C,B,A,B,C",
		"L,8,R,12,R,12,R,10",
		"R,10,R,12,R,10",
		"L,10,R,10,L,6",
	}
	for _, str := range instr {
		for _, s := range str {
			in <- int(s)
		}
		in <- '\n'
	}
	in <- 'n' // debug off
	in <- '\n'

	go func() {
		computer.Run()
		close(out)
	}()

	res := 0
	for i := range out {
		if i > 127 {
			res = i
		}
	}
	return res
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
