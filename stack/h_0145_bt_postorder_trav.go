package stack

import (
	"github.com/xunsavior/xunleetcode/tree"
)

/*
Company: Google, Amazon, Facebook, Uber

Given a binary tree, return the postorder traversal of its nodes' values.

Example:
Input: [1,null,2,3]
   1
    \
     2
    /
   3
Output: [3,2,1]

Follow up: Recursive solution is trivial, could you do it iteratively?
*/

/*
	We can use REVERSE postorder to store the node into stack, which
	is root -> right -> left
*/
func postorderTraversal145(root *tree.TNode) []int {
	if root == nil {
		return []int{}
	}
	output := []int{}
	stack := []*tree.TNode{root}
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		output = append([]int{current.Val}, output...)
		stack = stack[:len(stack)-1]

		l, r := current.Left, current.Right
		// note: it is stack that LIFO, so we store LEFT node first and then right node
		if l != nil {
			stack = append(stack, l)
		}
		if r != nil {
			stack = append(stack, r)
		}
	}
	return output
}
