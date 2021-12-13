package main

var testPuzzles = []struct {
	Puzzle    string
	Solution1 string
	Solution2 string
}{
	{
		Puzzle: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
		Solution1: "10",
		Solution2: "36",
	},
	{
		Puzzle: `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
		Solution1: "19",
		Solution2: "103",
	},
	{
		Puzzle: `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
		Solution1: "226",
		Solution2: "3509",
	},
	{
		Puzzle: puzzle,
		Solution1: "5228",
		Solution2: "131228",
	},
}

var puzzle = `EO-jc
end-tm
jy-FI
ek-EO
mg-ek
jc-jy
FI-start
jy-mg
mg-FI
jc-tm
end-EO
ds-EO
jy-start
tm-EO
mg-jc
ek-jc
tm-ek
FI-jc
jy-EO
ek-jy
ek-LT
start-mg`
