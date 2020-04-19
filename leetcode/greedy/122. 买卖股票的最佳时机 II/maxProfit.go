package main

import "math"

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	profit := 0
	price := math.MaxInt64
	for i:=0; i<len(prices); i++ {
		if prices[i] > price {
			profit += prices[i] - price
		}
		price = prices[i]
	}
	return profit
}
