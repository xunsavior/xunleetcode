package tree

/*
Company: Adobe(3)

Given a binary search tree and the lowest and highest boundaries as L and R,
trim the tree so that all its elements lies in [L, R] (R >= L). You might need
to change the root of the tree, so the result should return the new root of the
trimmed binary search tree.

Example 1:
Input:
    1
   / \
  0   2

  L = 1
  R = 2
Output:
    1
      \
	   2

Example 2:
Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3
Output:
      3
     /
   2
  /
 1

*/

func trimBST669(root *TNode, L int, R int) *TNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil && root.Val >= L && root.Val <= R {
		return root
	}

	if root.Val > R {
		return trimBST669(root.Left, L, R)
	}

	if root.Val < L {
		return trimBST669(root.Right, L, R)
	}

	return &TNode{
		Val:   root.Val,
		Left:  trimBST669(root.Left, L, R),
		Right: trimBST669(root.Right, L, R),
	}
}
