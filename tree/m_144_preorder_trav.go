package tree

/*
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

func preorderTraversal144(root *TNode) []int {
	if root == nil {
		return nil
	}
	nums := []int{}
	helper144(root, &nums)
	return nums
}

func helper144(current *TNode, nums *[]int) {
	if current == nil {
		return
	}
	*nums = append(*nums, current.Val)
	helper144(current.Left, nums)
	helper144(current.Right, nums)
}
