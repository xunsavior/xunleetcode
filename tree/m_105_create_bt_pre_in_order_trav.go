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
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	return buildTreeHelper(preorder, inorder, 0, 0, len(inorder)-1)
}

func buildTreeHelper(preOrder, inOrder []int, preStart, inStart, inEnd int) *TNode {
	if inStart > inEnd {
		return nil
	}
	currentNode := &TNode{
		Val:   preOrder[preStart],
		Left:  nil,
		Right: nil,
	}
	// find the index of the value in inOrder slice
	tempInStart := inStart
	for i := inStart; i <= inEnd; i++ {
		if inOrder[i] == currentNode.Val {
			tempInStart = i
			break
		}
	}
	currentNode.Left = buildTreeHelper(preOrder, inOrder, preStart+1, inStart, tempInStart-1)
	currentNode.Right = buildTreeHelper(preOrder, inOrder, preStart+tempInStart-inStart+1, tempInStart+1, inEnd)
	return currentNode
}
