package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"image"
	"strconv"
	"sync"
	"time"
)

func setup(input string) map[int]intcode.Computer {
	computers := make(map[int]intcode.Computer, 50)
	for i := 0; i < 50; i++ {
		in := make(chan int)
		out := make(chan int)
		c, err := intcode.Init(input, in, out)
		if err != nil {
			panic(err)
		}
		computers[i] = c
		go func() {
			c.Run()
			close(out)
		}()
		in <- i // send the address of this computer
	}
	return computers
}

func solve(input string) (string, string) {
	p1 := 0
	p2 := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		p1 = part1(setup(input))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		p2 = part2(setup(input))
		wg.Done()
	}()
	wg.Wait()

	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func part2(network map[int]intcode.Computer) int {
	nat := image.Pt(0, 0)
	nathist := image.Pt(-1, -1)
	queue := make(map[int][]image.Point, len(network))
	for {
		idle := true
		for addr, c := range network {
			p := image.Pt(-1, -1)
			if len(queue[addr]) > 0 {
				p = queue[addr][0]
			}
			select {
			case t := <-c.Out:
				p := image.Pt(<-c.Out, <-c.Out)
				if t == 255 {
					nat = p
				} else {
					queue[t] = append(queue[t], p)
				}
				idle = false
			case c.In <- p.X:
				if p.X != -1 {
					c.In <- p.Y
					queue[addr] = queue[addr][1:]
					idle = false
				}
			}
		}
		if !idle {
			continue
		}
		for _, q := range queue {
			if len(q) > 0 {
				idle = false
				break
			}
		}
		if idle {
			if nat == nathist {
				return nat.Y
			}
			queue[0] = append(queue[0], nat)
			nathist = nat
		}
	}
}

func part1(network map[int]intcode.Computer) int {
	for {
		for _, c := range network {
			select {
			case t := <-c.Out:
				x := <-c.Out
				y := <-c.Out
				if t == 255 {
					return y
				}
				network[t].In <- x
				network[t].In <- y
			case c.In <- -1:
			}
		}
	}
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
