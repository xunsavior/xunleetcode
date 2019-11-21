package leetcodetest

import (
	"leetcode/tree"
	"log"
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

type isSameTreeInput struct {
	node1  *tree.TNode
	node2  *tree.TNode
	expRes bool
}

var isSameTreeInputCases = []*isSameTreeInput{
	&isSameTreeInput{
		node1: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		node2: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		expRes: true,
	},
	&isSameTreeInput{
		node1: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		node2: &tree.TNode{
			Val:  1,
			Left: nil,
			Right: &tree.TNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		expRes: false,
	},
	&isSameTreeInput{
		node1: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		node2: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		expRes: false,
	},
}

func TestIsSameTree(t *testing.T) {
	for _, v := range isSameTreeInputCases {
		log.Printf("%v %v", *v.node1, *v.node2)
		equals(t, v.expRes, tree.IsSameTree(v.node1, v.node2))
	}
}

type isSymmetricInput struct {
	input  *tree.TNode
	expRes bool
}

var isSymmetricInputCases = []*isSymmetricInput{
	&isSymmetricInput{
		input: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val: 2,
				Left: &tree.TNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &tree.TNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &tree.TNode{
				Val: 2,
				Left: &tree.TNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &tree.TNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		expRes: true,
	},
	&isSymmetricInput{
		input: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:  2,
				Left: nil,
				Right: &tree.TNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &tree.TNode{
				Val:  2,
				Left: nil,
				Right: &tree.TNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		expRes: false,
	},
	&isSymmetricInput{
		input: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val: 2,
				Left: &tree.TNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &tree.TNode{
				Val: 2,
				Left: &tree.TNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		expRes: false,
	},
	&isSymmetricInput{
		input: &tree.TNode{
			Val: 1,
			Left: &tree.TNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		expRes: false,
	},
	&isSymmetricInput{
		input: &tree.TNode{
			Val: 5,
			Left: &tree.TNode{
				Val:  4,
				Left: nil,
				Right: &tree.TNode{
					Val: 1,
					Left: &tree.TNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			Right: &tree.TNode{
				Val:  1,
				Left: nil,
				Right: &tree.TNode{
					Val: 4,
					Left: &tree.TNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		expRes: false,
	},
}

func TestIsSymmetric(t *testing.T) {
	for i, v := range isSymmetricInputCases {
		log.Printf("%d %v %v", i, v.expRes, tree.IsSymmetric(v.input))
		equals(t, v.expRes, tree.IsSymmetric(v.input))
	}
}
