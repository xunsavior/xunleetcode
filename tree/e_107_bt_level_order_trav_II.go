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

func levelOrderBottom(root *TNode) [][]int {
	if root == nil {
		return nil
	}
	output := [][]int{}
	helper107(root, 0, &output)
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output
}

//top-down solution
func helper107(current *TNode, depth int, output *[][]int) {
	if current == nil {
		return
	}

	if depth > len(*output)-1 {
		*output = append(*output, []int{current.Val})
	} else {
		(*output)[depth] = append((*output)[depth], current.Val)
	}

	helper107(current.Left, depth+1, output)
	helper107(current.Right, depth+1, output)
}
