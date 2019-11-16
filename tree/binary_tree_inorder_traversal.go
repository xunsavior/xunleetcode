package tree

/*
	Given a binary tree, return the inorder traversal of its nodes' values.
	Example:
	Input: [1,null,2,3]
	1
	 \
	  2
	 /
	3
	Output: [1,3,2]
*/

// InOrderTraversal ...
func InOrderTraversal(root *TNode) []int {
	output := make([]int, 0)
	return addValToSlice(root, output)
}

func addValToSlice(currentNode *TNode, output []int) []int {
	// Special case: return empty slice if root is nil
	if currentNode == nil {
		return []int{}
	}

	lNode := currentNode.Left
	rNode := currentNode.Right

	/*
		if the LEFT node of CURRENT node is NOT nil, we simply move to that
		LEFT node
	*/
	if lNode != nil {
		output = addValToSlice(lNode, output)
	}

	/*
		if the LEFT node of current node is nil, we simply add the value of
		current node to the slice
	*/
	output = append(output, currentNode.Val)

	/*
		if the RIGHT node of current node is NOT nil, we simply move to that
		RIGHT node
	*/
	if rNode != nil {
		output = addValToSlice(rNode, output)
	}

	return output
}
