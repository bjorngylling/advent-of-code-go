package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type World map[int]image.Rectangle

func contestedPoints(w World) map[image.Point]struct{} {
	result := make(map[image.Point]struct{})
	for id, r := range w {
		for otherId, other := range w {
			if id == otherId {
				continue
			}
			intersect := r.Intersect(other)
			if !intersect.Empty() {
				for x := intersect.Min.X; x < intersect.Max.X; x++ {
					for y := intersect.Min.Y; y < intersect.Max.Y; y++ {
						result[image.Pt(x, y)] = struct{}{}
					}
				}
			}
		}
	}
	return result
}

func uncontestedClaims(w World) (result []int) {
	contestedIds := make(map[int]struct{})
	for id, r := range w {
		for otherId, other := range w {
			if id == otherId {
				continue
			}
			if r.Overlaps(other) {
				contestedIds[id] = struct{}{}
				contestedIds[otherId] = struct{}{}
			}
		}
	}
	for id := range w {
		if _, ok := contestedIds[id]; !ok {
			result = append(result, id)
		}
	}
	return
}

func createWorld(s []string) World {
	w := make(World)
	for _, ln := range s {
		strs := regexp.MustCompile(`\d+`).FindAllString(ln, -1)
		var ints []int
		for _, s := range strs {
			if i, err := strconv.Atoi(s); err == nil {
				ints = append(ints, i)
			}
		}
		w[ints[0]] = image.Rect(ints[1], ints[2], ints[1]+ints[3], ints[2]+ints[4])
	}
	return w
}

func main() {
	fileContent, err := ioutil.ReadFile("03_no_matter_how_you_slice_it/day3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	world := createWorld(strings.Split(string(fileContent), "\n"))

	fmt.Printf("Day 3 part 1 result: %d\n", len(contestedPoints(world)))

	fmt.Printf("Day 3 part 2 result: %v\n", uncontestedClaims(world)[0])
}
