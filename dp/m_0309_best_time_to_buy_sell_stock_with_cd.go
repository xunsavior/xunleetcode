package dp

import "math"

/*
Company: Google(2); Amazon(2), Apple(2)

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions
as you like (ie, buy one and sell one share of the stock multiple times) with the
following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
Example:

Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*/

type subtask struct {
	canBuy  bool
	canSell bool
	index   int
}

func maxProfit309(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	cache := make(map[subtask]int)
	return helper209(true, false, 0, prices, cache)
}

func helper209(canBuy, canSell bool, index int, prices []int, cache map[subtask]int) int {
	if index == len(prices) {
		return 0
	}

	val := prices[index]
	key := subtask{
		canSell: canSell,
		canBuy:  canBuy,
		index:   index,
	}
	if res, ok := cache[key]; ok {
		return res
	}

	if !canBuy && !canSell {
		res := helper209(true, false, index+1, prices, cache)
		cache[key] = res
		return res
	}

	if canBuy == true {
		buyRes := -val + helper209(false, true, index+1, prices, cache)
		passRes := helper209(canBuy, canSell, index+1, prices, cache)
		res := int(math.Max(float64(buyRes), float64(passRes)))
		cache[key] = res
		return res
	}

	sellRes := val + helper209(false, false, index+1, prices, cache)
	passRes := helper209(canBuy, canSell, index+1, prices, cache)
	res := int(math.Max(float64(sellRes), float64(passRes)))
	cache[key] = res
	return res
}
