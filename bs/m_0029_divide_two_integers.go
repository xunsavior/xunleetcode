package bs

import "math"

/*
Company: Facebook(21), Amazon(3), Riot Games(2), Apple(2), Bloomberg(2)

Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.

Example 1:
Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = truncate(3.33333..) = 3.

Example 2:
Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = truncate(-2.33333..) = -2.

Note:
Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed
integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns
2^31 − 1 when the division result overflows.
*/

func divide29(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	posDividend, posDivisor := int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	res := 0

	for posDividend >= posDivisor {
		bit := 0
		for posDividend >= (posDivisor << 1 << bit) {
			bit++
		}
		res += (1 << bit)
		posDividend -= (posDivisor << bit) // handle the remaining part
	}

	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		return res
	}

	return -res
}
