package main

//https://leetcode-cn.com/problems/domino-and-tromino-tiling/

//思路参考：https://leetcode-cn.com/problems/domino-and-tromino-tiling/solution/zhao-gui-lu-by-ccnuacmhdu/
//dp[i] = 2*dp[i-1] + dp[i-3]
func numTilings(N int) int {
	if N == 1 {return 1}
	if N == 2 {return 2}
	if N == 3 {return 5}
	dp1, dp2, dp3 := 1, 2, 5
	for i:=4; i<=N; i++ {
		dp3,dp2,dp1 = (2*dp3 + dp1)%1000000007,dp3, dp2
	}
	return dp3
}
