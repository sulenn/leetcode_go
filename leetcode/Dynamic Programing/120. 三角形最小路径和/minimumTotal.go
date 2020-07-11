package main

//https://leetcode-cn.com/problems/triangle/

//自底向上，dp[j] = max(triangle[i][j]+dp[j], triangle[i][j]+dp[j+1])
//其中，dp[j] 和 dp[j+1] 是 triangle[i][j] 第 i 行第 j 个元素的下一行 i+1 行相邻的两个元素
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = max(triangle[i][j]+dp[j], triangle[i][j]+dp[j+1])
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return b
	}
	return a
}
