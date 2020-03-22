package backtracking

import "sort"

/*
Company: Amazon(2); Microsoft(2); Bloomberg(3)

Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set MUST NOT contain duplicate subsets.

Example:
Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/

// Time: O(2^(n+1))
// Space: O(n)
func subsetsWithDup90(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	helper90(0, nums, []int{}, &res)
	return res
}

func helper90(index int, nums, addedNums []int, res *[][]int) {
	if len(nums) == index {
		*res = append(*res, addedNums)
		return
	}

	newNums := append(addedNums, nums[index])
	helper90(index+1, nums, newNums[:len(newNums):len(newNums)], res)

	// important
	if len(addedNums) > 0 && addedNums[len(addedNums)-1] == nums[index] {
		return
	}

	helper90(index+1, nums, addedNums, res)
}
