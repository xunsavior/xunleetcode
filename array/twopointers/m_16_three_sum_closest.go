package twopointers

import (
	"math"
	"sort"
)

/*
Company: Google(4), Amazon(2), Bloomberg(2), Apple(3)

Given an array nums of n integers and an integer target, find three integers in nums
such that the sum is closest to target. Return the sum of the three integers. You
may assume that each input would have exactly one solution.

Example:
Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest16(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[n-1]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			val := nums[i] + nums[l] + nums[r]
			valDistance := int(math.Abs(float64(target) - float64(val)))
			resDistance := int(math.Abs(float64(target) - float64(res)))
			if valDistance < resDistance {
				res = val
			}
			if val == target {
				return target
			} else if val < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
