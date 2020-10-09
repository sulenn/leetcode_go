package main

//https://leetcode-cn.com/problems/unique-paths-ii/submissions/

//动态规划：dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	r_len := len(obstacleGrid)
	c_len := len(obstacleGrid[0])
	dp := make([][]int, r_len)
	for i := 0; i < r_len; i++ {
		dp[i] = make([]int, c_len)
	}
	dp[0][0] = 1
	for i := 1; i < c_len; i++ { // 第一行赋值
		if dp[0][i-1] == 0 || obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}
	for i := 1; i < r_len; i++ { // 第一列赋值
		if dp[i-1][0] == 0 || obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}
	for i := 1; i < r_len; i++ {
		for j := 1; j < c_len; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[r_len-1][c_len-1]
}
