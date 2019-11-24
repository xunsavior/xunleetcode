package tree

/*
	Given inorder and postorder traversal of a tree, construct the binary tree.

	Note:
	You may assume that duplicates do not exist in the tree.

	For example, given
	inorder = [9,3,15,20,7]
	postorder = [9,15,7,20,3]
	Return the following binary tree:

		3
	   / \
	  9  20
		/  \
	   15   7
*/

/*
	Solution
	we can find the ROOT in postorder slice (from end to start)
	we find the index of that ROOT in inorder slice, and split it with left and right nodes
*/
func buildTree106(inorder, postorder []int) *TNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	return buildTreeHelper106(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
}

func buildTreeHelper106(inorder, postorder []int, pstart, istart, iend int) *TNode {
	if pstart < 0 || istart > iend {
		return nil
	}
	node := &TNode{
		Val:   postorder[pstart],
		Left:  nil,
		Right: nil,
	}

	tempInStart := istart
	for i := istart; i <= iend; i++ {
		if inorder[i] == node.Val {
			tempInStart = i
			break
		}
	}

	node.Left = buildTreeHelper106(inorder, postorder, pstart-1-(iend-tempInStart), istart, tempInStart-1)
	node.Right = buildTreeHelper106(inorder, postorder, pstart-1, tempInStart+1, iend)

	return node
}
