package tree

import "math"

/*
Company: Facebook

Also see: 124, 687

Given a binary tree, you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.
*/

func diameterOfBinaryTree543(root *TNode) int {
	maxPathSum := 0
	helper543(root, &maxPathSum)
	return maxPathSum
}

func helper543(current *TNode, maxPathSum *int) int {
	/*
	   we add 1 in advance regardless of the nullability of right and left node,
	   therefore we need to return -1 instead of 0 to eliminate that
	*/
	if current == nil {
		return -1
	}

	l, r := current.Left, current.Right

	lv, rv := helper543(l, maxPathSum)+1, helper543(r, maxPathSum)+1
	sum := lv + rv
	if sum > *maxPathSum {
		*maxPathSum = sum
	}
	return int(math.Max(float64(lv), float64(rv)))
}
