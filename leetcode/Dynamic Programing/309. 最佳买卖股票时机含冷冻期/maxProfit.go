package main

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	if len(prices) == 2 {
		if prices[1] > prices[0] {
			return prices[1] - prices[0]
		} else {
			return 0
		}
	}
	dp1 := 0
	dp2 := 0
	if prices[1] > prices[0] {
		dp2 = prices[1] - prices[0]
	}
	for i := 2; i < len(prices); i++ {
		if prices[i]-prices[i-2]+dp1 > dp2 {
			dp2, dp1 = prices[i]-prices[i-2]+dp1, dp2
		}
	}
	return dp2
}
