package bs

import "math"

/*
Company: Atlassian(6), Amazon(5), Facebook(3), Apple(3), Oracle(3)

Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
Follow up: Could you improve it to O(n log n) time complexity?
*/

func lengthOfLIS300(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	for _, num := range nums {
		start, end := 0, res
		for start < end {
			mid := start + (end-start)/2
			if dp[mid] < num {
				start = mid + 1
			} else {
				end = mid
			}
		}
		dp[start] = num
		if start == res {
			res++
		}
	}
	return res
}

// O(n^2)
func lengthOfLIS300WithDP(nums []int) int {
	res, dp := 0, make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(helper300(i, i, nums, dp)), float64(res)))
	}

	return res
}

func helper300(lastIndex, index int, nums []int, dp [][]int) int {
	if index == len(nums) {
		return 0
	}

	res, currentVal := 0, nums[index]

	if res = dp[lastIndex][index]; res != 0 {
		return res
	}

	if currentVal > nums[lastIndex] {
		res = int(math.Max(float64(helper300(lastIndex, index+1, nums, dp)), float64(1+helper300(index, index+1, nums, dp))))
	} else {
		if lastIndex == index {
			res = 1 + helper300(lastIndex, index+1, nums, dp)
		} else {
			res = helper300(lastIndex, index+1, nums, dp)
		}
	}

	dp[lastIndex][index] = res

	return res
}
