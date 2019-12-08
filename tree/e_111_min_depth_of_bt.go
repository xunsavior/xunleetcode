package tree

import "math"

/*
Also see: 104

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest
path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its minimum depth = 2.
*/

/*
	Important: If one child is nil, then we will return the min depth of the other child
*/
func minDepth111(root *TNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth111(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth111(root.Left) + 1
	}

	l, r := minDepth111(root.Left)+1, minDepth111(root.Right)+1
	return int(math.Min(float64(l), float64(r)))
}
