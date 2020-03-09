package tree

import "math"

/*
Company: Apple(2); Amazon(3); Google(6)

Also see: 124 and 543

Given a binary tree, find the length of the longest path where each node in the path has the same value.
This path may or may not pass through the root.
The length of path between two nodes is represented by the number of edges between them.

Example 1:
Input:

              5
             / \
            4   5
           / \   \
          1   1   5
Output: 2

Example 2:
Input:

              1
             / \
            4   5
           / \   \
          4   4   5
Output: 2
Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.

*/

func longestUnivaluePath687(root *TNode) int {
	if root == nil {
		return 0
	}
	maxPathSum := 0
	helper687(root, &maxPathSum)
	return maxPathSum
}

func helper687(current *TNode, maxPathSum *int) int {
	l, r := current.Left, current.Right
	if l == nil && r == nil {
		return 0 // we return 0, because we only count edge instead of node
	}

	lv, rv := 0, 0
	if l != nil {
		if l.Val == current.Val {
			lv = helper687(l, maxPathSum) + 1
		} else {
			// important: remember to reset right value to 0
			helper687(l, maxPathSum)
			lv = 0
		}
	}

	if r != nil {
		if r.Val == current.Val {
			rv = helper687(r, maxPathSum) + 1
		} else {
			// important: remember to reset right value to 0
			helper687(r, maxPathSum)
			rv = 0
		}
	}

	sum := lv + rv
	if sum > *maxPathSum {
		*maxPathSum = sum
	}

	return int(math.Max(float64(lv), float64(rv)))
}
