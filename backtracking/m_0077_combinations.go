package backtracking

/*
Company: Google(2), Microsoft(2); Amazon(2), Apple(2), Adobe(2)

Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine77(n int, k int) [][]int {
	res := [][]int{}
	helper77(1, n, k, []int{}, &res)
	return res
}

func helper77(index, n, k int, existed []int, res *[][]int) {
	if k == 0 {
		*res = append(*res, existed)
		return
	}

	// important: if n-index is even less than the number of slot, we directly return
	if n-index < k-1 {
		return
	}

	helper77(index+1, n, k, existed, res)
	newExisted := append(existed, index)
	helper77(index+1, n, k-1, newExisted[:len(newExisted):len(newExisted)], res)
}
