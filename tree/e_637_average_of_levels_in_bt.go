package tree

import (
	"leetcode/utils"
)

/*
Company: Facebook
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

Example 1:
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]

Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

Note:
The range of node's value is in the range of 32-bit signed integer.
*/

func averageOfLevels637(root *TNode) []float64 {
	nums := [][]float64{}
	helper637(root, 0, &nums)
	output := []float64{}
	for _, v := range nums {
		average := utils.Sum(v) / float64(len(v))
		output = append(output, average)
	}
	return output
}

// preorder traversal
func helper637(current *TNode, level int, nums *[][]float64) {
	if current == nil {
		return
	}

	if level > len(*nums)-1 {
		*nums = append(*nums, []float64{float64(current.Val)})
	} else {
		(*nums)[level] = append((*nums)[level], float64(current.Val))
	}

	helper637(current.Left, level+1, nums)
	helper637(current.Right, level+1, nums)
}
