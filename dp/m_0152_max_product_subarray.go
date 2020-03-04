package dp

/*
Given an integer array nums, find the contiguous subarray within an array
(containing at least one number) which has the largest product.

Example 1:
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largex`st product 6.

Example 2:
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

func maxProduct152(nums []int) int {
	res, max, min := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		val := nums[i]
		if val < 0 {
			max, min = min, max
		}

		if val > max*val {
			max = val
		} else {
			max *= val
		}

		if val < min*val {
			min = val
		} else {
			min *= val
		}

		if max > res {
			res = max
		}
	}

	return res
}
