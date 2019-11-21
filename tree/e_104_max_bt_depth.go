package tree

import "math"

/*
	Given a binary tree, find its maximum depth. The maximum depth is the number of nodes
	along the longest path from the root node down to the farthest leaf node.
	Note: A leaf is a node with no children.

	Example:
	Given binary tree [3,9,20,null,null,15,7],
	 3
	/ \
	9  20
      /  \
	 15   7
	return its depth = 3.
*/

// MaxDepth ...
func MaxDepth(root *TNode) int {
	if root == nil {
		return 0
	}
	return getMaxDepth(root, 1)
}

func getMaxDepth(currentNode *TNode, level int) int {
	// if current node has no child, then we will return depth
	if currentNode.Left == nil && currentNode.Right == nil {
		return level
	}

	// increase the depth
	level++

	// one of the chilren is nil
	if currentNode.Left == nil {
		return getMaxDepth(currentNode.Right, level)
	}
	if currentNode.Right == nil {
		return getMaxDepth(currentNode.Left, level)
	}

	// both children are NOT nil
	l, r := getMaxDepth(currentNode.Left, level), getMaxDepth(currentNode.Right, level)
	return int(math.Max(float64(l), float64(r)))
}
