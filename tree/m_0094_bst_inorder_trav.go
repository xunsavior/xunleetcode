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
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

// iterative way
func inorderTraversal94(root *TNode) []int {
	if root == nil {
		return nil
	}
	output := []int{}
	stack := []*TNode{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		l, r := top.Left, top.Right
		if l != nil {
			stack = append(stack, l)
		} else {
			output = append(output, top.Val)
			stack = stack[0 : len(stack)-1]
			// remember to set the left node of processed top node to nil, to avoid infinite loop
			if len(stack) > 0 {
				stack[len(stack)-1].Left = nil
			}
			if r != nil {
				stack = append(stack, r)
			}
		}
	}
	return output
}

// recoursion way
func helper94(current *TNode, output *[]int) {
	if current == nil {
		return
	}
	helper94(current.Left, output)
	*output = append(*output, current.Val)
	helper94(current.Right, output)
}
