package tree

import "math"

/*
Company: Zenefits

Given an array of numbers, verify whether it is the correct preorder
traversal sequence of a binary search tree. You may assume each number
in the sequence is unique.

Consider the following binary search tree:
     5
    / \
   2   6
  / \
 1   3
Example 1:
Input: [5,2,6,1,3]
Output: false

Example 2:
Input: [5,2,1,3,6]
Output: true

Follow up:
Could you do it using only constant space complexity?
*/
func verifyPreorder255(preorder []int) bool {
	nums := []int{}
	min := math.MinInt32
	for _, v := range preorder {
		if v < min {
			return false
		}
		for len(nums) > 0 && v > nums[len(nums)-1] {
			min = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
		nums = append(nums, v)
	}
	return true
}
