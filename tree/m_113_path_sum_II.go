package tree

/*
	Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

	Note: A leaf is a node with no children.

	Example:
	Given the below binary tree and sum = 22,
		5
	   / \
	  4   8
	 /   / \
    11  13  4
   /  \    / \
  7    2  5   1

  Return:
	[
	[5,4,11,2],
	[5,8,4,5]
	]
*/

// PathSum113 ...
func PathSum113(root *TNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	output := [][]int{}
	helper113(root, sum, []int{}, &output)
	return output
}

func helper113(current *TNode, sum int, nums []int, output *[][]int) {
	l, r := current.Left, current.Right
	nums = append(nums, current.Val)
	if l == nil && r == nil {
		if current.Val == sum {
			*output = append(*output, nums[:len(nums):len(nums)])
		}
		return
	}

	if r == nil || l == nil {
		if r == nil {
			helper113(l, sum-current.Val, nums[:len(nums):len(nums)], output)
		} else {
			helper113(r, sum-current.Val, nums[:len(nums):len(nums)], output)
		}
		return
	}

	helper113(l, sum-current.Val, nums[:len(nums):len(nums)], output)
	helper113(r, sum-current.Val, nums[:len(nums):len(nums)], output)
}
