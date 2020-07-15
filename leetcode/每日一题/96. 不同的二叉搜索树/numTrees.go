package main

//https://leetcode-cn.com/problems/unique-binary-search-trees/

//动态规划，dp[i] = dp[0]*dp[i-1] + dp[1]*dp[i-2] + ... + dp[i-2]*dp[1] + dp[i-1]*dp[0]
func numTrees(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
