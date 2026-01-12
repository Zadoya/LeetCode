// условие задачи - https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

package main

import "fmt"

func maxProfit(prices []int) int {
    sumProfit := 0

	buy := 0
	sell := 0
    for sell < len(prices) {
		switch {
		case prices[buy] <= prices[sell] && sell < len(prices) - 1 && prices[sell] < prices[sell+1]:
			sell++
		case prices[buy] > prices[sell]:
			buy = sell
			sell++
		default:
			sumProfit += prices[sell] - prices[buy]
			buy = sell
			sell++
		}
	}

	return sumProfit
}

const (
	s1 = "Привет, Надя!"
	s2 = "Привет! Дима"
)

func main() {
	fmt.Println(s1,s2)
	array := []int{1,2,3,4,5}
	fmt.Println(maxProfit(array))
}
