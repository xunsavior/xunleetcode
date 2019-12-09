package tree

/*
Also see: 102, 199

Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

func levelOrderBottom107(root *TNode) [][]int {
	items := [][]int{}
	// we create a map with k->[depth] and v->[slice per each row]
	itemMap := map[int][]int{}
	//use recursion to store the each row
	helper107(root, 0, itemMap)
	//move the read the slice from map using depth index
	for i := len(itemMap) - 1; i >= 0; i-- {
		items = append(items, itemMap[i])
	}
	return items
}

//StoreToItemMap ...
func helper107(currentNode *TNode, index int, itemMap map[int][]int) {
	if currentNode == nil {
		return
	}
	// store the value into the map
	itemMap[index] = append(itemMap[index], currentNode.Val)
	//increase depth by 1
	index++
	//store the next depth of node from LEFT to RIGHT
	helper107(currentNode.Left, index, itemMap)
	helper107(currentNode.Right, index, itemMap)
}
