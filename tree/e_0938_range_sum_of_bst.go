package tree

/*
Company: Facebook(14), Apple(3), Oracle(3), Amazon(2); Microsoft(3)

Given the root node of a binary search tree, return the sum
of values of all nodes with value between L and R (inclusive).
The binary search tree is guaranteed to have unique values.

Example 1:
Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32

Example 2:
Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23

Note:
The number of nodes in the tree is at most 10000.
The final answer is guaranteed to be less than 2^31.
*/

func rangeSumBST938(root *TNode, L int, R int) int {
	if root == nil {
		return 0
	}

	val := root.Val

	if val > R {
		return rangeSumBST938(root.Left, L, R)
	}

	if val < L {
		return rangeSumBST938(root.Right, L, R)
	}

	return val + rangeSumBST938(root.Left, L, R) + rangeSumBST938(root.Right, L, R)
}
