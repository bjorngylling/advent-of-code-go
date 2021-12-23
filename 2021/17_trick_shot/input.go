package main

var testPuzzles = []struct {
	Puzzle    string
	Solution1 string
	Solution2 string
}{
	{
		Puzzle:    "target area: x=20..30, y=-10..-5",
		Solution1: "45",
		Solution2: "112",
	},
	{
		Puzzle:    puzzle,
		Solution1: "6903",
		Solution2: "2351",
	},
}

var puzzle = `target area: x=235..259, y=-118..-62`
