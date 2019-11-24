package tree

/*
	Given a binary tree, return the zigzag level order traversal of its nodes' values.
	i.e. from left to right, then right to left for the next level and alternate between

	For example:
	Given binary tree [3,9,20,null,null,15,7],
		3
	   / \
	  9  20
		/  \
	   15   7
	return its zigzag level order traversal as:
	[
	[3],
	[20,9],
	[15,7]
	]
*/

//ZigzagLevelOrder ...
func ZigzagLevelOrder(root *TNode) [][]int {
	slices := [][]int{}
	itemMap := map[int][]int{}
	storeValToMap(root, 0, true, itemMap)
	for i := 0; i < len(itemMap); i++ {
		slices = append(slices, itemMap[i])
	}
	return slices
}

func storeValToMap(currentNode *TNode, depth int, isLeftToRight bool, itemMap map[int][]int) {
	if currentNode == nil {
		return
	}
	/*
		if it is from LEFT to RIGHT, we will directly append
		if it is from RIGHT to LEFT, we will prepend
	*/
	if isLeftToRight {
		itemMap[depth] = append(itemMap[depth], currentNode.Val)
	} else {
		itemMap[depth] = append([]int{currentNode.Val}, itemMap[depth]...)
	}
	// next layer and store it to the map
	storeValToMap(currentNode.Left, depth+1, !isLeftToRight, itemMap)
	storeValToMap(currentNode.Right, depth+1, !isLeftToRight, itemMap)
}
