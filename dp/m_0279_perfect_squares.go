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

f(1) = 1 -> 1
f(2) = f(1) + f(1) -> 2
f(3) = f(2) + f(1) -> 3
f(4) = f(4) -> 1
f(5) = f(4) + f(1) -> 2
f(6) = f(5) + f(4) -> 3
f(7) = f(6) + f(4) -> 4
f(8) = f(4) + f(4) -> 2
f(9) = f(9) -> 1
f(10) = f(9) + f(1) -> 2
f(11) = f(10) + f(9) -> 3
f(12) = f(4) + f(4) + f(4) -> 3
f(13) = f(4) + f(9) -> 2

*/

func numSquares279(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	maxSquareIndex := int(math.Sqrt(float64(n))) + 1
	nums := make([]int, maxSquareIndex)
	for i := 1; i < maxSquareIndex; i++ {
		nums[i] = i * i
	}

	for i := 1; i <= n; i++ {
		for s := 1; s < maxSquareIndex; s++ {
			if i < nums[s] {
				break
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-nums[s]]+1)))
		}
	}

	return dp[n]
}
