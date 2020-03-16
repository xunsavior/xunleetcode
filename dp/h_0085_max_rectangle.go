package dp

import "math"

/*
Company: Amazon(4), Google(3), Uber(3), Microsoft(2), Samsung(2)

Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

Example:

Input:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
Output: 6
*/

func maximalRectangle85(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	res := 0

	n := len(matrix[0])
	dp := make([]int, n)

	for i := range matrix {
		for j := range matrix[i] {
			// if the above element is 1, then we sum it up with current element; otherwise make current element 0
			if i == 0 {
				if matrix[i][j] == 49 {
					dp[j] = 1
				}
			} else {
				if matrix[i][j] != 48 {
					dp[j] = 1 + dp[j]
				} else {
					dp[j] = 0
				}
			}
		}
		res = int(math.Max(float64(res), float64(helper85(dp))))
	}

	return res
}

func helper85(heights []int) int {
	res, stack := 0, []int{-1}

	for i := range heights {
		for stack[len(stack)-1] != -1 && heights[i] <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := heights[pop] * (i - 1 - stack[len(stack)-1])
			res = int(math.Max(float64(area), float64(res)))
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area := heights[pop] * (len(heights) - 1 - stack[len(stack)-1])
		res = int(math.Max(float64(area), float64(res)))
	}

	return res
}
