package leetcodetest

import (
	"leetcode/tree"
	"testing"
)

type inOrderTraversal struct {
	input     *tree.TNode
	expectRes []int
}

var inOrderTraversalTestCases = []*inOrderTraversal{
	&inOrderTraversal{
		input: &tree.TNode{
			Val:  1,
			Left: nil,
			Right: &tree.TNode{
				Val: 2,
				Left: &tree.TNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		expectRes: []int{1, 3, 2},
	},
	&inOrderTraversal{
		input:     nil,
		expectRes: []int{},
	},
}

func TestInOrderTraversal(t *testing.T) {
	for _, v := range inOrderTraversalTestCases {
		equals(t, v.expectRes, tree.InOrderTraversal(v.input))
	}
}

type numTreesTestCase struct {
	input  int
	expRes int
}

var numTreesTestCases = []*numTreesTestCase{
	&numTreesTestCase{
		input:  3,
		expRes: 5,
	},
	&numTreesTestCase{
		input:  4,
		expRes: 14,
	},
}

func TestNumTrees(t *testing.T) {
	for _, v := range numTreesTestCases {
		equals(t, v.expRes, tree.NumTreesWithDPSolution(v.input))
	}
}

// 98. Validate Binary Search Tree
type validBSTInput struct {
	input *tree.TNode
	exRes bool
}

var validBSTCases = []*validBSTInput{
	&validBSTInput{
		&tree.TNode{
			Val: 2,
			Left: &tree.TNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		true,
	},
	&validBSTInput{
		&tree.TNode{
			Val: 5,
			Left: &tree.TNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val: 4,
				Left: &tree.TNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &tree.TNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
		},
		false,
	},
	&validBSTInput{
		&tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		false,
	},
	&validBSTInput{
		&tree.TNode{
			Val: 10,
			Left: &tree.TNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val: 15,
				Left: &tree.TNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &tree.TNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
			},
		},
		false,
	},
	&validBSTInput{
		&tree.TNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		true,
	},
}

func TestIsValidBST(t *testing.T) {
	for _, v := range validBSTCases {
		equals(t, v.exRes, tree.IsValidBST(v.input))
	}
}
