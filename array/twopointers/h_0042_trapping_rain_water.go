package twopointers

import "math"

/*
Company: Amazon(42), Microsoft(24), Facebook(12), Google(11), Apple(7)

Given n non-negative integers representing an elevation map where
the width of each bar is 1, compute how much water it is able to
trap after raining.

The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
In this case, 6 units of rain water (blue section) are being trapped.
Thanks Marcos for contributing this image!

Example:
Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
*/

func trap42(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0

	for l < r {
		if lMax < rMax {
			lMax = int(math.Max(float64(height[l]), float64(lMax)))
			res += lMax - height[l]
			l++
		} else {
			rMax = int(math.Max(float64(height[r]), float64(rMax)))
			res += rMax - height[r]
			r--
		}
	}

	return res
}
