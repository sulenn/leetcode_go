package main

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

//动态规划：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/go-jie-ti-dong-tai-gui-hua-shi-jian-fu-za-du-on-ko/
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	var (
		buy      = -prices[0]
		sell     = 0
		cooldown = 0
	)

	for i := 1; i < len(prices); i++ {
		buy, sell, cooldown = Max(buy, cooldown-prices[i]), Max(buy+prices[i], sell), sell
	}
	return Max(sell, cooldown)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
