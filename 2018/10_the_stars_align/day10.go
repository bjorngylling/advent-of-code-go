package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Star struct {
	Pos   image.Point
	Speed image.Point
}

func NewStar(x, y, dx, dy int) *Star {
	return &Star{image.Pt(x, y), image.Pt(dx, dy)}
}

var logLineRegexp = regexp.MustCompile(`<(.\d+), (.\d+)>`)

func parseStars(ln string) (r []*Star) {
	m := logLineRegexp.FindAllStringSubmatch(ln, -1)
	for i := 0; i < len(m); i += 2 {
		x, y := atoiPair(m[i][1], m[i][2])
		dx, dy := atoiPair(m[i+1][1], m[i+1][2])
		r = append(r, NewStar(x, y, dx, dy))
	}
	return
}

func atoiPair(a, b string) (i, j int) {
	i, _ = strconv.Atoi(strings.TrimSpace(a))
	j, _ = strconv.Atoi(strings.TrimSpace(b))
	return
}

func step(l []*Star) {
	for _, s := range l {
		s.Pos = s.Pos.Add(s.Speed)
	}
}

func findBounds(l []*Star) image.Rectangle {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := 0, 0
	for _, s := range l {
		if minX > s.Pos.X {
			minX = s.Pos.X
		}
		if maxX < s.Pos.X {
			maxX = s.Pos.X
		}
		if minY > s.Pos.Y {
			minY = s.Pos.Y
		}
		if maxY < s.Pos.Y {
			maxY = s.Pos.Y
		}
	}
	return image.Rect(minX, minY, maxX, maxY)
}

func createImage(l []*Star) *image.RGBA {
	// figure out the image bounds
	b := findBounds(l)
	img := image.NewRGBA(image.Rectangle{Min: b.Min.Sub(image.Pt(10, 10)), Max: b.Max.Add(image.Pt(10, 10))})

	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			img.Set(x, y, color.Black)
		}
	}

	for _, s := range l {
		img.Set(s.Pos.X, s.Pos.Y, color.White)
	}

	return img
}

func saveImage(name string, img *image.RGBA) {
	// Save to out.png
	f, _ := os.OpenFile(fmt.Sprintf("10_the_stars_align/%s.png", name), os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func main() {
	fileContent, err := ioutil.ReadFile("10_the_stars_align/day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	stars := parseStars(string(fileContent))

	for i := 1; i < 52000; i++ {
		step(stars)
		bounds := findBounds(stars)
		if bounds.Dx() < 300 && bounds.Dy() < 200 {
			saveImage(strconv.Itoa(i), createImage(stars))
		}
	}
}
