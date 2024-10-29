// условие задачи - https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		switch {
		case prices[i] < minPrice:
			minPrice = prices[i]
		case prices[i] - minPrice > max:
			max = prices[i] - minPrice
		}
	}

	return max
}