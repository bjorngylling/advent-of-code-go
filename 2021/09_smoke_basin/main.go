package main

import (
	"fmt"
	"image"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bjorngylling/advent-of-code/util"
)

var dirs = []image.Point{
	image.Pt(0, -1),
	image.Pt(1, 0),
	image.Pt(0, 1),
	image.Pt(-1, 0),
}

func solve(input string) (string, string) {
	m := map[image.Point]int{}
	for y, ln := range strings.Split(input, "\n") {
		for x, c := range ln {
			m[image.Pt(x, y)] = util.GetInt(string(c))
		}
	}

	risk := part1(m)
	p1 := 0
	for _, v := range risk {
		p1 += v
	}
	return strconv.Itoa(p1), strconv.Itoa(part2(m, risk))
}

func part1(m map[image.Point]int) map[image.Point]int {
	risk := map[image.Point]int{}
	for pt, l := range m {
		if n, ok := m[pt.Add(dirs[0])]; !ok || l < n {
			if e, ok := m[pt.Add(dirs[1])]; !ok || l < e {
				if s, ok := m[pt.Add(dirs[2])]; !ok || l < s {
					if w, ok := m[pt.Add(dirs[3])]; !ok || l < w {
						risk[pt] = l + 1
					}
				}
			}
		}
	}
	return risk
}

func part2(m map[image.Point]int, low map[image.Point]int) int {
	var basinSizes []int
	for pt := range low {
		b := flood(m, pt)
		basinSizes = append(basinSizes, len(b))
	}
	sort.Ints(basinSizes)
	l := len(basinSizes)
	return basinSizes[l-1] * basinSizes[l-2] * basinSizes[l-3]
}

func flood(m map[image.Point]int, start image.Point) map[image.Point]struct{} {
	basin := map[image.Point]struct{}{}
	valid := func(p image.Point) bool {
		l, ok := m[p]
		_, v := basin[p]
		return !v && ok && l < 9
	}

	var f func(p image.Point)
	f = func(p image.Point) {
		if !valid(p) {
			return
		}
		basin[p] = struct{}{}
		for _, d := range dirs {
			f(p.Add(d))
		}
	}
	f(start)

	return basin
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
