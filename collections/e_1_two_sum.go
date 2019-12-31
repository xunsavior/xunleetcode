package collections

/*
Company: Amazon(197), Google(88), Adobe(56), Apple(43), Facebook(26)
Given an array of integers, return INDICES of the two numbers such that they add up to a specific target.
You may assume that each input would have EXACTLY one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum1(nums []int, target int) []int {
	numMap := map[int]int{}
	for i, v := range nums {
		numMap[v] = i
	}

	for i, v := range nums {
		// remember to check if index != i
		if index, ok := numMap[target-v]; ok && index != i {
			return []int{i, index}
		}
	}

	return nil
}
