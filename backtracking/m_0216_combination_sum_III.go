package backtracking

/*
Company: nil; Amazon(2); Microsoft(2)

Find all possible combinations of k numbers that add up to a number n, given
that only numbers from 1 to 9 can be used and each combination should be a
unique set of numbers.

Note:
All numbers will be positive integers.
The solution set must not contain duplicate combinations.

Example 1:
Input: k = 3, n = 7
Output: [[1,2,4]]

Example 2:
Input: k = 3, n = 9
Output: [[1,2,6], [1,3,5], [2,3,4]]
*/

func combinationSum216(k int, n int) [][]int {
	res := [][]int{}
	helper216(0, 0, k, n, []int{}, &res)
	return res
}

func helper216(index, currentSum, k, n int, tmp []int, res *[][]int) {
	if k == len(tmp) || currentSum == n {
		if k == len(tmp) && currentSum == n {
			*res = append(*res, tmp)
		}
		return
	}

	if currentSum > n || index == 9 {
		return
	}

	for i := index + 1; i <= 9; i++ {
		newTmp := append(tmp, i)
		helper216(i, currentSum+i, k, n, newTmp[:len(newTmp):len(newTmp)], res)
	}
}
