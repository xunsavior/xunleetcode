package dp

/*
Company: Amazon(6), Apple(4), Microsoft(4), Facebook(4), Mathworks(4)

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach
the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right

Example 2:
Input: m = 7, n = 3
Output: 28
*/

func uniquePaths62(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	// initiate an slice
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	return helper62(0, 0, cache)
}

func helper62(m, n int, cache [][]int) int {
	if m == len(cache)-1 && n == len(cache[0])-1 {
		return 1
	}

	if m > len(cache)-1 || n > len(cache[0])-1 {
		return 0
	}

	if val := cache[m][n]; val != 0 {
		return val
	}

	res := helper62(m+1, n, cache) + helper62(m, n+1, cache)
	cache[m][n] = res
	return res
}
