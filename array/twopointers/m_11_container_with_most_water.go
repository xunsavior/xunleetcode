package twopointers

import "math"

/*
Company: Amazon(17), Apple(4), Yahoon(4), Google(3), Facebook(3)

Given n non-negative integers a1, a2, ..., an , where each represents a point
at coordinate (i, ai). n vertical lines are drawn such that the two endpoints
of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis
forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

func maxArea11(height []int) int {
	res := 0
	n := len(height)
	l, r := 0, n-1

	for l < r {
		lv, rv, width := height[l], height[r], r-l
		area := int(math.Min(float64(lv), float64(rv))) * width
		if area > res {
			res = area
		}
		if lv < rv {
			l++
		} else if rv < lv {
			r--
		} else {
			l++
			r--
		}
	}

	return res
}
