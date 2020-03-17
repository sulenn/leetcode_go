package main

import "math"

//https://leetcode-cn.com/problems/perfect-squares/
// 递归 CS-NOTES
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i:=1; i<=n;i++ {
		min := math.MaxInt64
		for j:=1; j*j <= i; j++ {
			if dp[i-j*j]+1 < min {
				min = dp[i-j*j]+1
			}
		}
		dp[i] = min
	}
	return dp[n]
}
