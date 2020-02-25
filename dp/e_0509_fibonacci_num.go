package dp

/*
Company: Amazon(5), Barclays(5), Microsoft(4), Apple(4), Facebook(2)

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
Fibonacci sequence, such that each number is the sum of the two preceding
ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(N) = F(N - 1) + F(N - 2), for N > 1.
Given N, calculate F(N).

Example 1:
Input: 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

Example 2:
Input: 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

Example 3:
Input: 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

Note:
0 ≤ N ≤ 30.
*/

func fib509(N int) int {
	if N == 0 {
		return 0
	}
	cache := make([]int, N-1)
	return helper509(N, cache)
}

func helper509(n int, cache []int) int {
	if n < 2 {
		return n
	}

	if cache[n-2] != 0 {
		return cache[n-2]
	}

	res := helper509(n-1, cache) + helper509(n-2, cache)
	cache[n-2] = res
	return res
}
