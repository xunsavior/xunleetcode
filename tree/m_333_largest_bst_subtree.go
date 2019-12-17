package tree

import "math"

/*
Company: Microsoft
Given a binary tree, find the largest subtree which is a Binary Search Tree (BST),
where largest means subtree with largest number of nodes in it.

Note:
A subtree must include all of its descendants.

Example:
Input: [10,5,15,1,8,null,7]
   10
   / \
  5  15
 / \   \
1   8   7
Output: 3

Explanation:
The Largest BST Subtree in this case is the highlighted one.
The return value is the subtree's size, which is 3.

Follow up:
Can you figure out ways to solve it with O(n) time complexity?
*/

func largestBSTSubtree333(root *TNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	helper333(root, &max)
	return max
}

/*
	Bottom-up solution

	Each layer will return 3 values:
	1. the size of current valid BST
	2. min value of the upper node
	3. max value of the upper node

	Note: min and max can be equal
*/
func helper333(current *TNode, maxSize *int) (size, min, max int) {
	// nil scenario
	if current == nil {
		size, min, max = 0, 0, 0
		return
	}

	l, r := current.Left, current.Right
	lsize, lmin, lmax := helper333(l, maxSize)
	rsize, rmin, rmax := helper333(r, maxSize)

	// invalid scenario
	if lsize == -1 || (lsize > 0 && lmax >= current.Val) || rsize == -1 || (rsize > 0 && rmin <= current.Val) {
		size, min, max = -1, 0, 0
		return
	}

	// valid scenario
	sum := lsize + rsize + 1
	if sum > *maxSize {
		*maxSize = sum
	}

	size = sum

	if min = current.Val; lsize != 0 {
		min = lmin
	}

	if max = current.Val; rsize != 0 {
		max = rmax
	}

	return
}
