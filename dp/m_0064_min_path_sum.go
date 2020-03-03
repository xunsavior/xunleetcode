package dp

import "math"

/*
Company: Amazon(15), Google(7), Goldman Sachs(5), Bloomberg(2), Microsoft(2)

Given a m x n grid filled with non-negative numbers, find a path from top
left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:
Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

func minPathSum64(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	return helper64(0, 0, grid, cache)
}

func helper64(m, n int, grid [][]int, cache [][]int) int {
	if m == len(grid)-1 && n == len(grid[0])-1 {
		return grid[m][n]
	}

	// important: we should avoid adding those cases going wide or long
	if m == len(grid) || n == len(grid[0]) {
		return math.MaxInt32
	}

	if res := cache[m][n]; res != 0 {
		return res
	}

	downCost, rightCost := grid[m][n]+helper64(m+1, n, grid, cache), grid[m][n]+helper64(n, n+1, grid, cache)
	res := int(math.Min(float64(downCost), float64(rightCost)))
	cache[m][n] = res
	return res
}
