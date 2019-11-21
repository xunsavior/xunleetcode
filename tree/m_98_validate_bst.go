package tree

/*
	Given a binary tree, determine if it is a valid binary search tree (BST).

	Assume a BST is defined as follows:
	The left subtree of a node contains only nodes with keys less than the node's key.
	The right subtree of a node contains only nodes with keys greater than the node's key.
	Both the left and right subtrees must also be binary search trees.
*/

// IsValidBST ... Solution 1: limit the value range for the sub-trees
func IsValidBST(root *TNode) bool {
	return checkIfIsValid(root, nil, nil)
}

func checkIfIsValid(currentNode *TNode, min, max *int) bool {
	// when we reach the leaf node
	if currentNode == nil {
		return true
	}
	// if the value of current node is NOT int the range of (min, max), return false
	if (max != nil && currentNode.Val >= *max) ||
		(min != nil && currentNode.Val <= *min) {
		return false
	}

	return checkIfIsValid(currentNode.Left, min, &(currentNode.Val)) && checkIfIsValid(currentNode.Right, &(currentNode.Val), max)
}

// Solution 2: using inorder traversel and check if it is sorted or not
