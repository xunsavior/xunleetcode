package backtracking

/*
Company: Airbnb(9), Amazon(8), Facebook(5), Microsoft(3), Google(2)

Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.

Example 1:
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]

Example 2:
Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	res := [][]int{}
	helper39(0, target, candidates, []int{}, &res)
	return res
}

func helper39(index, remainder int, candidates, tmp []int, res *[][]int) {
	if remainder == 0 {
		*res = append(*res, tmp)
		return
	}

	if remainder < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		temp := append(tmp, candidates[i])
		n := len(temp)
		helper39(i, remainder-candidates[i], candidates, temp[0:n:n], res)
	}

}
