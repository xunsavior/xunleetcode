package tree

/*
	Given a binary tree, determine if it is a valid binary search tree (BST).

	Assume a BST is defined as follows:
	The left subtree of a node contains only nodes with keys less than the node's key.
	The right subtree of a node contains only nodes with keys greater than the node's key.
	Both the left and right subtrees must also be binary search trees.
*/

// IsValidBST ...
func IsValidBST(root *TNode) bool {
	if root == nil {
		return true
	}
	return checkIfIsValid(nil, nil, root, nil)
}

func checkIfIsValid(min, max *int, currentNode *TNode, parentNode *TNode) bool {
	// when we reach the leaf node
	if currentNode == nil {
		return true
	}
	// we must check if the value of current node is in (min, max)
	if (min != nil && currentNode.Val <= *min) || (max != nil && currentNode.Val >= *max) {
		return false
	}

	isLeftValid, isRightValid := true, true
	lNode, rNode := currentNode.Left, currentNode.Right
	// check the left part
	if lNode != nil {
		if parentNode != nil {
			/*
				important:
				not only we need to check if left node is less than current node,
				we should also pay attention to the entire tree. (Same as the right side)
				e.g. [3,1,5,0,2,4,6,null,null,null,3] and [3,null,30,10,null,null,15,null,45] are invalid

				That is why we need both parent and current node to set up the max and min
				value for the next child node.
			*/
			if parentNode.Right == currentNode {
				isLeftValid = checkIfIsValid(&(parentNode.Val), &(currentNode.Val), lNode, currentNode)
			} else {
				isLeftValid = checkIfIsValid(min, &(currentNode.Val), lNode, currentNode)
			}
		} else {
			isLeftValid = checkIfIsValid(min, &(currentNode.Val), lNode, currentNode)
		}
	}
	//check the right part
	if rNode != nil {
		if parentNode != nil {
			if parentNode.Left == currentNode {
				isRightValid = checkIfIsValid(&(currentNode.Val), &(parentNode.Val), rNode, currentNode)
			} else {
				isRightValid = checkIfIsValid(&(currentNode.Val), max, rNode, currentNode)
			}
		} else {
			isRightValid = checkIfIsValid(&(currentNode.Val), max, rNode, currentNode)
		}
	}
	// we need validate if left and right side are both valid
	return isRightValid && isLeftValid
}
