package day7

import "testing"

func TestCreateTree(t *testing.T) {
	tree := CreateTree(sampleInput)

	if tree.Root.Name != "tknk" {
		t.Errorf("Expected root node to be tknk, but was %s instead", tree.Root.Name)
	}
	if len(tree.Root.Children) != 3 {
		t.Errorf("Expected root node to have 3 children, but had %d instead", len(tree.Root.Children))
	}
}

func TestSumWeight(t *testing.T) {
	tree := CreateTree(sampleInput)

	if weight := sumWeight(tree.nodeTable["ugml"]); weight != 251 {
		t.Errorf("Expected the total weight of node ugml to be 251, but was %d instead", weight)
	}
}

func TestCalculateChildWeights(t *testing.T) {
	expected := map[string]int {"ugml": 251, "padx": 243, "fwft": 243}
	tree := CreateTree(sampleInput)

	weights := calculateWeights(tree.Root.Children)
	if l := len(weights); l != 3 {
		t.Errorf("Expected there to be 3 children, but was %d instead", l)
	}
	for _, e := range weights {
		if e.weight != expected[e.node.Name] {
			t.Errorf("Expected the total weight of node %s to be %d, but was %d instead",
				e.node.Name, expected[e.node.Name], e.weight)
		}
	}
}

func TestHeaviest(t *testing.T) {
	tree := CreateTree(sampleInput)

	if n := heaviest(calculateWeights(tree.Root.Children)); n.node.Name != "ugml" {
		t.Errorf("Expected the heaviest child to be ugml, but was %s instead", n.node.Name)
	}
}

func TestLightest(t *testing.T) {
	tree := CreateTree(sampleInput)

	if n := lightest(calculateWeights(tree.Root.Children)); n.node.Name != "padx" {
		t.Errorf("Expected the heaviest child to be ugml, but was %s instead", n.node.Name)
	}
}

func TestFindOverweightNode(t *testing.T) {
	tree := CreateTree(sampleInput)

	node, weightDifference := FindOverweightNode(tree.Root)
	if node.Name != "ugml" {
		t.Errorf("Expected overweight node to be ugml, but was %s instead", node.Name)
	}
	if weightDifference != 8 {
		t.Errorf("Expected overweight node to be 8 units too heavy, but was %d instead", weightDifference)
	}
}

const sampleInput = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`