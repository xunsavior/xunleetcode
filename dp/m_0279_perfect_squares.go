package dp

import "math"

/*
Company: Uber(4), Microsoft(2), Cisco(2); Google(5), Apple(3)

Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.

Example 2:
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.

*/

// numSquares(n)= min(numSquares(n-k) + 1) ∀k ∈ {square numbers}
func numSquares279(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	maxSquareIndex := int(math.Sqrt(float64(n))) + 1
	squareNums := make([]int, maxSquareIndex)
	for i := 1; i < maxSquareIndex; i++ {
		squareNums[i] = i * i
	}

	for i := 1; i <= n; i++ {
		for s := 1; s < maxSquareIndex; s++ {
			if i < squareNums[s] {
				break
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-squareNums[s]]+1)))
		}
	}

	return dp[n]
}
