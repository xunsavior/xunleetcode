package tree

/*
	Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

	Note: A leaf is a node with no children.

	Example:
	Given the below binary tree and sum = 22,
		5
	   / \
	  4   8
	 /   / \
    11  13  4
	/  \      \
   7    2      1
	return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/

//Solution: we add current val to the previous sum
func hasPathSum112(root *TNode, sum int) bool {
	if root == nil {
		return false
	}
	return helper112(root, sum, 0)
}

func helper112(currentNode *TNode, sum int, parentSum int) bool {
	l, r := currentNode.Left, currentNode.Right
	// current node is leaf
	if l == nil && r == nil {
		return sum == parentSum+currentNode.Val
	}
	// current node ONLY has left child
	if r == nil {
		return helper112(l, sum, parentSum+currentNode.Val)
	}
	// current node ONLY has right child
	if l == nil {
		return helper112(r, sum, parentSum+currentNode.Val)
	}
	// current node has both left and right child
	return helper112(l, sum, parentSum+currentNode.Val) ||
		helper112(r, sum, parentSum+currentNode.Val)
}
