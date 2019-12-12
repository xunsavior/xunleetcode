package tree

/*
Given a complete binary tree, count the number of nodes.

Note:
Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level, except possibly the last,
is completely filled, and all nodes in the last level are as far left as possible.
It can have between 1 and 2h nodes inclusive at the last level h.

Example:
Input:
    1
   / \
  2   3
 / \  /
4  5 6
Output: 6
*/

func countNodes222(root *TNode) int {
	if root == nil {
		return 0
	}
	return helper222(root) + 1
}

// bottom-up solution
func helper222(current *TNode) int {
	if current == nil {
		return 0
	}
	l, r := current.Left, current.Right
	ln, rn := 0, 0
	if l != nil {
		ln = helper222(l) + 1
	}

	if r != nil {
		rn = helper222(r) + 1
	}
	return ln + rn
}
