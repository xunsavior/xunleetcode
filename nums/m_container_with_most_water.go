package nums

import (
	"log"
	"math"
)

/*
	https://leetcode.com/problems/container-with-most-water/
*/

func maxArea(height []int) int {
	//get the value of first and last element
	lv, rv := height[0], height[len(height)-1]

	// select the min height
	var h int
	if h = lv; rv < lv {
		h = rv
	}

	// calculate current area value
	// Note: the width = len(height) - 1
	currentS := h * (len(height) - 1)

	// find the base case, in which there are only 2 elements (width = 1)
	if len(height) == 2 {
		return currentS
	}

	/*
		important:
		we move from left to right or right to left by 1 unit based on whose value is smaller
	*/
	if lv < rv {
		// left -> right by 1 unit
		return int(math.Max(float64(maxArea(height[1:])), float64(currentS)))
	}
	// right -> left by 1 unit
	return int(math.Max(float64(maxArea(height[:len(height)-1])), float64(currentS)))
}

// TestMaxArea ...
func TestMaxArea() {
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	log.Println(maxArea(height1))
}
