package dp

import "math"

/*
Company: ;; Amazon(3), Yahoo(2)

On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum
cost to reach the top of the floor, and you can either start from the step with index 0,
or the step with index 1.

Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.

Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].

Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
*/

func minCostClimbingStairs746(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	cache := make([]*int, len(cost))
	val0, val1 := helper746(0, cost, cache), helper746(1, cost, cache)
	return int(math.Min(float64(val0), float64(val1)))
}

func helper746(index int, cost []int, cache []*int) int {
	if index >= len(cost) {
		return 0
	}

	if res := cache[index]; res != nil {
		return *res
	}

	val1, val2 := cost[index]+helper746(index+1, cost, cache), cost[index]+helper746(index+2, cost, cache)
	res := int(math.Min(float64(val1), float64(val2)))
	cache[index] = &res
	return res
}
