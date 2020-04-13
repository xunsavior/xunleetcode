package bs

/*
Company: Facebook(16), Amazon(7), LinkedIn(4), Google(3), Bloomberg(3)

Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2^-2 = 1/(2*2) = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/

// Time: O(logN)
// Space: O(1)
func myPow50(x float64, n int) float64 {
	start, end, N := float64(1), x, n
	if n < 0 {
		end = 1 / end
		N = -n
	}

	for N != 0 {
		if N%2 == 1 {
			start *= end
		}
		end *= end
		N /= 2
	}

	return end
}
