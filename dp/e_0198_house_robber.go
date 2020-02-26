package dp

import "math"

/*
Company: Google(9), Amazon(3), Expedia(3), Cisco(3), Microsoft(2)

You are a professional robber planning to rob houses along a street. Each house has
a certain amount of money stashed, the only constraint stopping you from robbing each
of them is that adjacent houses have security system connected and it will automatically
contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
			 Total amount you can rob = 1 + 3 = 4.

Example 2:
Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/

func rob198(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cache := make([]*int, len(nums))
	return helper198(0, nums, cache)
}

func helper198(index int, nums []int, cache []*int) int {
	if index >= len(nums) {
		return 0
	}

	if res := cache[index]; res != nil {
		return *res
	}

	val1, val2 := nums[index]+helper198(index+2, nums, cache), helper198(index+1, nums, cache)
	res := int(math.Max(float64(val1), float64(val2)))
	cache[index] = &res
	return res
}
