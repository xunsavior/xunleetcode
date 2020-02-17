package dfsandbfs

/*
Company: Amazon(2); Microsoft(3), Adobe(3); Apple(2)

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

// BFS
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// we need to store all nodes in the next level
		nextLevel, vals := []*TreeNode{}, []int{}
		for _, node := range queue {
			vals = append(vals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		queue = nextLevel
		res = append([][]int{vals}, res...)
	}

	return res
}
