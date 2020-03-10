package dp

import "math"

/*
Company: Google(13), Amazon(8), Uber(7), Huawei(4), Bloomberg(3)

Given a 2D binary matrix filled with 0's and 1's, find the largest
square containing only 1's and return its area.

Example:
Input:
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0
Output: 4

*/

func maximalSquare221(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	maxDP := 0
	// create a dp matrix with all 0, and every element will be regarded as the bottom right corner vertex of squre
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// then we start from i=1, j=1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				// we know the two edges ought to be the exactly the same as the matrix
				if matrix[i][j] == 49 {
					dp[i][j] = 1
				}
			} else if matrix[i][j] == 49 {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), math.Min(float64(dp[i][j-1]), float64(dp[i-1][j-1])))) + 1
			}
			maxDP = int(math.Max(float64(dp[i][j]), float64(maxDP)))
		}
	}

	return maxDP * maxDP
}
