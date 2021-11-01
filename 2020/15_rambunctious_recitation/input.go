package main

var testPuzzles = []struct {
	Puzzle    string
	Solution1 string
	Solution2 string
}{
	{
		Puzzle:    "0,3,6",
		Solution1: "436",
		Solution2: "175594",
	},
	{
		Puzzle:    "1,3,2",
		Solution1: "1",
		Solution2: "2578",
	},
	{
		Puzzle:    "2,1,3",
		Solution1: "10",
		Solution2: "3544142",
	},
	{
		Puzzle:    "1,2,3",
		Solution1: "27",
		Solution2: "261214",
	},
	{
		Puzzle:    "2,3,1",
		Solution1: "78",
		Solution2: "6895259",
	},
	{
		Puzzle:    "3,2,1",
		Solution1: "438",
		Solution2: "18",
	},
	{
		Puzzle:    "3,1,2",
		Solution1: "1836",
		Solution2: "362",
	},
}

var puzzle = `8,13,1,0,18,9`
