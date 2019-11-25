package tree

/*
	Given preorder and inorder traversal of a tree, construct the binary tree.
	Note:
	You may assume that duplicates do not exist in the tree.

	For example, given
	preorder = [3,9,20,15,7]
	inorder = [9,3,15,20,7]
	Return the following binary tree:
		3
	   / \
	  9  20
		/  \
	   15   7
*/

/*
	Solution
	we can find the ROOT in preorder slice (from start to end)
	we find the index of that ROOT in inorder slice, and split it with left and right nodes
*/
func buildTree(preorder []int, inorder []int) *TNode {
	presize, insize := len(preorder), len(inorder)

	if presize == 0 || insize == 0 {
		return nil
	}

	// find the root index in inorder slice
	tmpStart := 0
	for i := 0; i < insize; i++ {
		if inorder[i] == preorder[0] {
			tmpStart = i
			break
		}
	}

	// we set preorder[0] as the root node, and get its left and right using recursion
	return &TNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:tmpStart+1], inorder[:tmpStart]),
		Right: buildTree(preorder[tmpStart+1:], inorder[tmpStart+1:]),
	}
}
