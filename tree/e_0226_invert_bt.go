package tree

/*
Company: Google(3), VMware(3), Amazon(2); Bloomberg(2); Microsoft(2)

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
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	return &TNode{
		Val:   root.Val,
		Left:  invertTree226(root.Left),
		Right: invertTree226(root.Right),
	}
}
