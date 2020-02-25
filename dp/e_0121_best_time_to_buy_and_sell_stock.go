package dp

import "math"

/*
Company: Amazon(33), Microsoft(14), Facebook(9), Google(5), Apple(3)

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell
one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
			 Not 7-1 = 6, as selling price needs to be larger than buying price.

Example 2:
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

/*
	Solution:
	The points of interest are the peaks and valleys in the given graph.
	We need to find the largest peak following the smallest valley.
*/
func maxProfit121(prices []int) int {
	minCost, maxProfit := math.MaxInt32, 0

	for _, v := range prices {
		if v < minCost {
			minCost = v
		} else if v-minCost > maxProfit {
			maxProfit = v - minCost
		}
	}

	return maxProfit
}
