package main

var testPuzzles = []struct {
	Puzzle    string
	Solution1 string
	Solution2 string
}{
	{
		Puzzle: `....#
#..#.
#..##
..#..
#....`,
		Solution1: "2129920",
		Solution2: "99",
	},
	{
		Puzzle:    puzzle,
		Solution1: "7543003",
		Solution2: "",
	},
}

var puzzle = `#....
#...#
##.##
....#
#.##.`
