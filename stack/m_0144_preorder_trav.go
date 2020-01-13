package stack

import "leetcode/tree"

/*
Company: Amazon(2); Cisco(2), Google(2);
Given a binary tree, return the preorder traversal of its nodes' values.

Example:
Input: [1,null,2,3]
   1
    \
     2
    /
   3
Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

// iterative approach
func preorderTraversal144(root *tree.TNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*tree.TNode{root}
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		l, r := pop.Left, pop.Right
		res = append(res, pop.Val)
		stack = stack[:len(stack)-1]
		if r != nil {
			stack = append(stack, r)
		}
		if l != nil {
			stack = append(stack, l)
		}
	}
	return res
}

// recursion
func helper144(current *tree.TNode, nums *[]int) {
	if current == nil {
		return
	}
	*nums = append(*nums, current.Val)
	helper144(current.Left, nums)
	helper144(current.Right, nums)
}
