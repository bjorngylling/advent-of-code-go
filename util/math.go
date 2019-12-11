package util

import (
	"image"
	"math"
	"strconv"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Return the smallest of all supplied ints
func Min(ints ...int) int {
	switch len(ints) {
	case 0:
		panic("no ints specified")
	case 1:
		return ints[0]
	case 2:
		return min(ints[0], ints[1])
	default:
		curMin := ints[0]
		for _, i := range ints[1:] {
			curMin = min(curMin, i)
		}
		return curMin
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Return the largest of all supplied ints
func Max(ints ...int) int {
	switch len(ints) {
	case 0:
		panic("no ints specified")
	case 1:
		return ints[0]
	case 2:
		return max(ints[0], ints[1])
	default:
		curMax := ints[0]
		for _, i := range ints[1:] {
			curMax = max(curMax, i)
		}
		return curMax
	}
}

// Abs returns the absolute value of the given number
func Abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// GetInt returns the given string as an int, or panics if it is invalid
func GetInt(in string) int {
	res, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return res
}

// RoundDown returns the given float rounded towards 0
func Trunc(in float64) int {
	return int(math.Trunc(in))
}

func ManhattanDistance(a, b image.Point) int {
	return Abs(b.X-a.X) + Abs(b.Y-a.Y)
}
