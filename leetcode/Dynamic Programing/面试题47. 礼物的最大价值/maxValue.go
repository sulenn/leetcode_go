package main

//https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

//dp[i][j] += max(dp[i-1][j],dp[i][j-1])
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	for i := 1; i < len(grid); i++ { // 先处理第一列， 避免后面溢出
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ { //同理处理第一行
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
