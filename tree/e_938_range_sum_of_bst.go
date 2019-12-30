package tree

/*
Company: Facebook(10), Amazon(2), Microsoft(2), Apple(2)

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
	l, r := root.Left, root.Right

	if l == nil && r == nil && root.Val >= L && root.Val <= R {
		return root.Val
	}

	if root.Val > R {
		return rangeSumBST938(l, L, R)
	}

	if root.Val < L {
		return rangeSumBST938(r, L, R)
	}

	return root.Val + rangeSumBST938(l, L, R) + rangeSumBST938(r, L, R)
}
