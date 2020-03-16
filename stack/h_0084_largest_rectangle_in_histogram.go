package stack

import "math"

/*
Company: Amazon(8), Microsoft(3), Bloomberg(3), Google(2); Flipkart(2)

Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
find the area of largest rectangle in the histogram.


Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
The largest rectangle is shown in the shaded area, which has area = 10 unit.

*/

func largestRectangleArea84(heights []int) int {
	res, stack := 0, []int{-1}

	for i := range heights {
		for stack[len(stack)-1] != -1 && heights[i] <= heights[stack[len(stack)-1]] {
			// in this case, we check 6, and then 5 * 2
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
