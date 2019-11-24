package tree

/*
	Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

	For example:
	Given binary tree [3,9,20,null,null,15,7],
	 3
	/ \
   9  20
	 /  \
	15   7
	return its level order traversal as:
	[
	[3],
	[9,20],
	[15,7]
	]
*/

func levelOrder(root *TNode) [][]int {
	items := [][]int{}
	if root == nil {
		return [][]int{}
	}
	// we create a map with k->[depth] and v->[slice per each row]
	itemMap := map[int][]int{}
	//use recursion to store the each row
	StoreToItemMap(root, 0, itemMap)
	//move the read the slice from map using depth index
	for i := 0; i < len(itemMap); i++ {
		items = append(items, itemMap[i])
	}
	return items
}
