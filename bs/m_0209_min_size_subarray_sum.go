package bs

import "math"

/*
Company: Goldman Sachs(18), Google(8), Facebook(4), Amazon(3), Oracle(3)

Given an array of n positive integers and a positive integer s, find the minimal
length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.

Follow up:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
*/

func minSubArrayLen209(s int, nums []int) int {
	res := 0

	if len(nums) == 0 {
		return res
	}

	if nums[0] >= s {
		return res + 1
	}

	i, j, sum := 0, 1, nums[0]
	for j < len(nums) {
		sum += nums[j]
		dif := sum - s
		if dif >= 0 {
			for dif >= nums[i] {
				sum -= nums[i]
				dif -= nums[i]
				i++
			}
			if res == 0 {
				res = j - i + 1
			} else {
				res = int(math.Min(float64(res), float64(j-i+1)))
			}
		}
		j++
	}

	return res
}
