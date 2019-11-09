package nums

/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.
	Example:
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1]
*/

func twoSumWithBruteForce(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumWithTwoPassMap(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i := range nums {
		complement := target - nums[i]
		i2, ok := m[complement]
		if ok && i2 != i {
			return []int{i, i2}
		}
	}
	return nil
}
