package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"sync"
	"time"
)

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func solve(input string) (string, string) {
	phaseSetting := []int{0, 1, 2, 3, 4}
	part1 := 0
	for p := make([]int, len(phaseSetting)); p[0] < len(p); nextPerm(p) {
		thrust := 0
		for _, i := range getPerm(phaseSetting, p) {
			out := make(chan int, 10)
			in := make(chan int, 10)
			in <- i
			in <- thrust
			computer, err := intcode.Init(input, in, out)
			if err != nil {
				panic(err)
			}
			computer.Run()
			thrust = <-out
		}
		part1 = util.Max(part1, thrust)
	}

	phaseSetting = []int{5, 6, 7, 8, 9}
	part2 := 0
	for p := make([]int, len(phaseSetting)); p[0] < len(p); nextPerm(p) {
		cfg := getPerm(phaseSetting, p)

		wg := &sync.WaitGroup{}
		wg.Add(5)

		amp1In := make(chan int, 10)
		amp1In <- cfg[0]
		amp1In <- 0
		amp1Out := make(chan int, 10)
		amp1Out <- cfg[1] // Prepare phase settings for amp 2
		amp1, err := intcode.Init(input, amp1In, amp1Out)
		if err != nil {
			panic(err)
		}
		go func() {
			amp1.Run()
			wg.Done()
		}()

		amp2Out := make(chan int, 10)
		amp2Out <- cfg[2] // Prepare phase settings for amp 3
		amp2, err := intcode.Init(input, amp1Out, amp2Out)
		if err != nil {
			panic(err)
		}
		go func() {
			amp2.Run()
			wg.Done()
		}()

		amp3Out := make(chan int, 10)
		amp3Out <- cfg[3] // Prepare phase settings for amp 4
		amp3, err := intcode.Init(input, amp2Out, amp3Out)
		if err != nil {
			panic(err)
		}
		go func() {
			amp3.Run()
			wg.Done()
		}()

		amp4Out := make(chan int, 10)
		amp4Out <- cfg[4] // Prepare phase settings for amp 5
		amp4, err := intcode.Init(input, amp3Out, amp4Out)
		if err != nil {
			panic(err)
		}
		go func() {
			amp4.Run()
			wg.Done()
		}()

		amp5, err := intcode.Init(input, amp4Out, amp1In)
		if err != nil {
			panic(err)
		}
		go func() {
			amp5.Run()
			wg.Done()
		}()

		wg.Wait() // wait for all programs to terminate
		part2 = util.Max(part2, <-amp1In)
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
