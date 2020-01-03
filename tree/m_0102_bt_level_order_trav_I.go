package tree

/*
Also see: 107, 199

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

func levelOrder102(root *TNode) [][]int {
	items := [][]int{}
	helper102(root, 0, &items)
	return items
}

/*
	Use preorder traversal to store every val to the 2d array
	from left to right
*/
func helper102(current *TNode, depth int, output *[][]int) {
	if current == nil {
		return
	}
	// we need to check if the array have already added this row(depth) of nodes
	if depth > len(*output)-1 {
		*output = append(*output, []int{current.Val})
	} else {
		(*output)[depth] = append((*output)[depth], current.Val)
	}

	helper102(current.Left, depth+1, output)
	helper102(current.Right, depth+1, output)
}
