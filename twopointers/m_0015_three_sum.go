package twopointers

import "sort"

/*
Company: Amazon(26), Facebook(17), Microsoft(9), Google(7), Apple(5)

Given an array nums of n integers, are there elements a, b, c in nums such
that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:
The solution set must not contain duplicate triplets.

Example:
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

// ThreeSum ...
func threeSum15(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n < 3 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// check if the value of current index is equal to the one of the last index
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			if tmp := nums[i] + nums[l] + nums[r]; tmp == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// remember to check if the value of next index is the same as the current one to avoid repeated cases
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if tmp < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
