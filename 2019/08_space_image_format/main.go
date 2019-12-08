package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	"time"
)

func solve(input string) (string, string) {
	layers := readLayers(input, 25, 6)

	img := image.NewGray(image.Rect(0, 0, 25, 6))
	for l := len(layers) - 1; l >= 0; l-- {
		for i, p := range layers[l] {
			x, y := i%25, i/25
			switch p {
			case 0:
				img.Set(x, y, color.Black)
			case 1:
				img.Set(x, y, color.White)
			}
		}
	}

	f, _ := os.Create("part2.png")
	defer f.Close()
	png.Encode(f, img)

	return strconv.Itoa(imgHash(layers)), "see part2.png"
}

func imgHash(img [][]int) int {
	hashLayer := 0
	zeroCount := math.MaxInt32
	for i, l := range img {
		curZeroCount := countPixels(l, 0)
		if curZeroCount < zeroCount {
			hashLayer = i
			zeroCount = curZeroCount
		}
	}
	return countPixels(img[hashLayer], 1) * countPixels(img[hashLayer], 2)
}

func countPixels(l []int, p int) int {
	c := 0
	for _, i := range l {
		if p == i {
			c++
		}
	}
	return c
}

func readLayers(input string, w int, h int) [][]int {
	layers := make([][]int, len(input)/(w*h))
	for i, c := range input {
		l := i / (w * h)
		layers[l] = append(layers[l], util.GetInt(string(c)))
	}
	return layers
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
