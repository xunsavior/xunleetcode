package tree

/*
	Company: Amazon(5), Google(3), Microsoft(2)
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

func inorderTraversal94(root *TNode) []int {
	output := []int{}
	helper94(root, &output)
	return output
}

func helper94(current *TNode, output *[]int) {
	if current == nil {
		return
	}
	helper94(current.Left, output)
	*output = append(*output, current.Val)
	helper94(current.Right, output)
}
