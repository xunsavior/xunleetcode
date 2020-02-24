package dp

import "math"

/*
Company: Twitter(19), LinkedIn(6); Apple(2);

There are a row of n houses, each house can be painted with one of the three colors: red,
blue or green. The cost of painting each house with a certain color is different. You have
to paint all the houses such that no two adjacent houses have the same color.

The cost of painting each house with a certain color is represented by a n x 3 cost matrix.
For example, costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is
the cost of painting house 1 with color green, and so on... Find the minimum cost to paint
all houses.

Note:
All costs are positive integers.

Example:

Input: [[17,2,17],[16,16,5],[14,3,19]]
Output: 10
Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue.
             Minimum cost: 2 + 5 + 3 = 10.
*/

const (
	red   = 0
	green = 1
	blue  = 2
)

func minCost256(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	cache := make([][3]int, n)
	redCost, greenCost, blueCost := costs[0][red], costs[0][green], costs[0][blue]
	res1, res2, res3 := redCost+helper256(red, 1, costs, cache), greenCost+helper256(green, 1, costs, cache), blueCost+helper256(blue, 1, costs, cache)
	return int(math.Min(float64(res1), math.Min(float64(res2), float64(res3))))
}

// f(color0, i) = val + min(f(color1, i+1), f(color2, i+1))
// color0 != color1 && color0 != color2
func helper256(lastColor, index int, costs [][]int, cache [][3]int) int {
	if index == len(costs) {
		return 0
	}

	res1, res2, res3 := cache[index][red], cache[index][green], cache[index][blue]
	redCost, greenCost, blueCost := costs[index][red], costs[index][green], costs[index][blue]

	if res1 == 0 {
		res1 = redCost + helper256(red, index+1, costs, cache)
		cache[index][red] = res1
	}

	if res2 == 0 {
		res2 = greenCost + helper256(green, index+1, costs, cache)
		cache[index][green] = res2
	}

	if res3 == 0 {
		res3 = blueCost + helper256(blue, index+1, costs, cache)
		cache[index][blue] = res3
	}

	if lastColor == red {
		return int(math.Min(float64(res2), float64(res3)))
	}

	if lastColor == green {
		return int(math.Min(float64(res1), float64(res3)))
	}

	return int(math.Min(float64(res1), float64(res2)))
}
