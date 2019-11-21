package tree

/*
	Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

	For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
		1
	   / \
	  2   2
	 / \ / \
	3  4 4  3

	But the following [1,2,2,null,3,null,3] is not:
		1
	   / \
	  2   2
	   \   \
	    3   3
	2, 3, 1, 2, 3

	[1,2,2,2,null,2] is not:
		1
	   / \
	  2   2
	 /   /
	2   2
	Note:
	Bonus points if you could solve it both recursively and iteratively.
*/

// IsSymmetric ...
func IsSymmetric(root *TNode) bool {
	if root == nil {
		return true
	}
	return checkIfSymmetric(root.Left, root.Right)
}

func checkIfSymmetric(l, r *TNode) bool {
	// if left and right node are both nil, then it is symmetric
	if l == nil && r == nil {
		return true
	}
	// if only one of them is nil, then it is NOT symmetric
	if l == nil || r == nil {
		return false
	}
	// if the values of two children are different, then it is NOT symmetric
	if l.Val != r.Val {
		return false
	}

	return checkIfSymmetric(l.Left, r.Right) && checkIfSymmetric(l.Right, r.Left)
}
