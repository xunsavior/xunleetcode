package dp

import "math"

/*
Company: Amazon(4); Apple(4), Google(2), Bloomberg(2); Houzz(2)

Given a triangle, find the minimum path sum from top to bottom.
Each step you may move to adjacent numbers on the row below.

For example, given the following triangle
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

*/

// Top down solution
func minimumTotal120TopDown(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	cache := make([][]int, len(triangle))
	for i := range cache {
		cache[i] = make([]int, len(triangle[i]))
	}

	return helper120(0, 0, triangle, cache)
}

func helper120(m, n int, triangle [][]int, cache [][]int) int {
	if m == len(triangle)-1 && n >= 0 && n <= len(triangle[m]) {
		return triangle[m][n]
	}

	if res := cache[m][n]; res != 0 {
		return res
	}

	val := triangle[m][n]
	left, right := val+helper120(m+1, n, triangle, cache), val+helper120(m+1, n+1, triangle, cache)
	res := int(math.Min(float64(left), float64(right)))
	cache[m][n] = res
	return res
}

// bottom up solution using stack
func minimumTotal120BottomUp(triangle [][]int) int {
	bottom := triangle[len(triangle)-1]
	stack := triangle[:len(triangle)-1]

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for i := range top {
			top[i] += int(math.Min(float64(bottom[i]), float64(bottom[i+1])))
		}

		bottom = top
	}

	return bottom[0]
}
