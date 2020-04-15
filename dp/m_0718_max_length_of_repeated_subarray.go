package dp

import "math"

/*
Company: Indeed(5); nil; Intuit(5), Pinterest(5), Citadel

Given two integer arrays A and B, return the maximum length of an subarray that appears in both arrays.

Example 1:
Input:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
Output: 3
Explanation:
The repeated subarray with maximum length is [3, 2, 1].

Note:
1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100
*/

// Time: O(MN)
// Space: O(MN)
func findLength718(A []int, B []int) int {
	res := 0

	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}

	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				res = int(math.Max(float64(res), float64(dp[i][j])))
			}
		}
	}

	return res
}
