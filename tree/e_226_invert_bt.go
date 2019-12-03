package tree

/*
Invert a binary tree.

Example:
Input:
     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:
     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/

func invertTree226(root *TNode) *TNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	root.Left, root.Right = invertTree226(root.Right), invertTree226(root.Left)
	return root
}
