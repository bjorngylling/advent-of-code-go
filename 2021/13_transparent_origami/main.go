package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bjorngylling/advent-of-code/util"
)

func fold(pnts map[image.Point]struct{}, fld image.Point) map[image.Point]struct{} {
	res := map[image.Point]struct{}{}

	for pt := range pnts {
		if pt.Y < fld.Y || pt.X < fld.X {
			res[pt] = struct{}{}
			continue
		}

		d := pt.Sub(fld).Mul(2)
		if fld.X == 0 {
			d.X = 0
		} else if fld.Y == 0 {
			d.Y = 0
		}
		res[pt.Sub(d)] = struct{}{}
	}
	return res
}

func solve(input string) (string, string) {
	spl := strings.Split(input, "\n\n")
	pnts := map[image.Point]struct{}{}
	for _, ln := range strings.Split(spl[0], "\n") {
		p := image.Point{}
		fmt.Sscanf(ln, "%d,%d", &p.X, &p.Y)
		pnts[p] = struct{}{}
	}
	var folds []image.Point
	for _, ln := range strings.Split(spl[1], "\n") {
		ln = strings.TrimPrefix(ln, "fold along ")
		s := strings.Split(ln, "=")
		p := image.Point{}
		if s[0] == "x" {
			p.X = util.GetInt(s[1])
		} else {
			p.Y = util.GetInt(s[1])
		}
		folds = append(folds, p)
	}

	pt2 := pnts
	for _, f := range folds {
		pt2 = fold(pt2, f)
	}
	img := image.NewGray(image.Rect(0, 0, 50, 10))
	for p := range pt2 {
		img.Set(p.X, p.Y, color.White)
	}
	f, _ := os.Create("part2.png")
	png.Encode(f, img)

	return strconv.Itoa(len(fold(pnts, folds[0]))), "see part2.png"
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
