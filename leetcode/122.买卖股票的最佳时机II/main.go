package main

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0
	cur := prices[0]
	for i:=1; i<len(prices); i++ {
		if prices[i] > cur {
			result += prices[i] - cur
		}
		cur = prices[i]
	}
	return result
}

func main() {
	
}