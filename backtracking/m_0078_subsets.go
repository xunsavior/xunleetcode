package backtracking

/*
Company: Facebook(10), Amazon(10), Bloomberg(3), Microsoft(2), Google(2)

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:
Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func subsets78(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	helper78(0, []int{}, nums, &res)
	return res
}

func helper78(index int, current, nums []int, res *[][]int) {
	if len(nums) == index {
		*res = append(*res, current)
		return
	}

	helper78(index+1, current, nums, res)
	newNums := append(current, nums[index])
	helper78(index+1, newNums[:len(newNums):len(newNums)], nums, res)
}
