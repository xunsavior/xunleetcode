package tree

import "math"

/*
Also see: 111

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the
longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

// MaxDepth104 ... bottom-up solution
func MaxDepth104(root *TNode) int {
	if root == nil {
		return 0
	}
	r, l := MaxDepth104(root.Left)+1, MaxDepth104(root.Right)+1
	return int(math.Max(float64(l), float64(r)))
}
