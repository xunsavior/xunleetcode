package tree

import "math"

/*
Also see: 543, 687

Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes
from some starting node to any node in the tree along the parent-child connections.
The path must contain at least one node and does not need to go through the root.

Example 1:
Input: [1,2,3]

       1
      / \
     2   3

Output: 6

Example 2:
Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7
Output: 42
*/

func maxPathSum124(root *TNode) int {
	maxSum := math.MinInt32
	helper124(root, &maxSum)
	return maxSum
}

func helper124(current *TNode, maxSum *int) int {
	if current == nil {
		return 0
	}
	l, r := current.Left, current.Right
	// calculate max value of left subtree and right subtree (remember to compare it with 0)
	lv, rv := int(math.Max(0, float64(helper124(l, maxSum)))), int(math.Max(0, float64(helper124(r, maxSum))))
	// calculate the sum of max left subtree, max right subtree and current value
	sum := lv + rv + current.Val
	// continue update the "global" maxSum by comparing it to the sum
	*maxSum = int(math.Max(float64(*maxSum), float64(sum)))
	/*
	   Note:
	   1. we must use root
	   2. at most ONE child can be used (because one path cannot be gone through twice)
	*/
	return current.Val + int(math.Max(float64(lv), float64(rv)))
}
