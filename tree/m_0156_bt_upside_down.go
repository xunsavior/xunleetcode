package tree

/*
Given a binary tree where all the right nodes are either leaf nodes
with a sibling (a left node that shares the same parent node) or empty,
flip it upside down and turn it into a tree where the original right
nodes turned into left leaf nodes. Return the new root.

Example:
Input: [1,2,3,4,5]
    1
   / \
  2   3
 / \
4   5

Output: return the root of the binary tree [4,5,2,#,#,3,1]
   4
  / \
 5   2
    / \
   3   1
*/

func upsideDownBinaryTree156(root *TNode) *TNode {
	if root == nil || root.Left == nil {
		return root
	}
	// we ONLY return the new root (the most left leaf node)
	newRoot := upsideDownBinaryTree156(root.Left)

	root.Left.Left, root.Left.Right = root.Right, root
	root.Left, root.Right = nil, nil

	return newRoot
}
