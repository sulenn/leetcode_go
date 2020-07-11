package main

//https://leetcode-cn.com/problems/minimum-cost-for-tickets/

//动态规划，参考：https://leetcode-cn.com/problems/minimum-cost-for-tickets/solution/li-jie-gai-ti-de-he-xin-bu-zai-dong-tai-gui-hua-by/
func mincostTickets(days []int, costs []int) int {
	maxDay := days[len(days)-1]
	dp := make([]int, maxDay+1) // dp 长度为最大天数+1
	day := 1
	day_1 := 0
	day_7 := 0
	day_30 := 0
	for i := 0; i < len(days); {
		if days[i] > day { // 当前不旅游
			dp[day] = dp[day-1]
		} else { // 当天旅游，dp[i] = min(dp[i - 1] + costs[0],dp_7 + costs[1], dp_30 + costs[2])，注意边界溢出
			day_1 = dp[day-1] + costs[0]
			if day-7 < 0 {
				day_7 = dp[0] + costs[1]
			} else {
				day_7 = dp[day-7] + costs[1]
			}
			if day-30 < 0 {
				day_30 = dp[0] + costs[2]
			} else {
				day_30 = dp[day-30] + costs[2]
			}
			dp[day] = min(day_1, day_7, day_30)
			i++
		}
		day++
	}
	return dp[maxDay]
}

func min(a, b, c int) int {
	if a >= b {
		if b >= c {
			return c
		} else {
			return b
		}
	} else {
		if a >= c {
			return c
		} else {
			return a
		}
	}
	return -1
}
