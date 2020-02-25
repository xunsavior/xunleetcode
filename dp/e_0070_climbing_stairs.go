package dp

/*
Company: Amazon(8), Uber(5), Adobe(5), Microsoft(5), Oracle(3)

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:
Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

*/

func climbStairs70(n int) int {
	cache := make([]int, n-1)
	return helper70(n, cache)
}

func helper70(n int, cache []int) int {
	if n < 2 {
		return 1
	}

	if cache[n-2] != 0 {
		return cache[n-2]
	}

	res := helper70(n-1, cache) + helper70(n-2, cache)
	cache[n-2] = res
	return res
}
