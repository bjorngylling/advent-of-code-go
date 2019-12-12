package main

var testPuzzles = []struct {
	Puzzle    string
	Steps     int
	Solution1 string
	Solution2 string
}{
	{
		Puzzle: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
		Steps:     10,
		Solution1: "179",
		Solution2: "2772",
	},
	{
		Puzzle: `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
		Steps:     100,
		Solution1: "1940",
		Solution2: "4686774924",
	},
	{
		Puzzle:    puzzle,
		Steps:     1000,
		Solution1: "9999",
		Solution2: "",
	},
}

var puzzle = `<x=14, y=9, z=14>
<x=9, y=11, z=6>
<x=-6, y=14, z=-4>
<x=4, y=-4, z=-3>`
