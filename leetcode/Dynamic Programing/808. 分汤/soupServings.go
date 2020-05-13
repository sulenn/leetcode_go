package main

//https://leetcode-cn.com/problems/soup-servings/

//参考：https://leetcode-cn.com/problems/soup-servings/solution/leetcode-fen-tang-dong-tai-gui-hua-ji-yi-sou-suo-b/

//dp[i][j] = 0.25*(dp[i-4][j]+dp[i-3][j-1]+dp[i-2][j-2]+dp[i-1][j-3])
func soupServings(N int) float64 {
	if N >= 4800 {return 1.0}
	times := N/25
	if N%25 > 0 {
		N = times + 1
	} else {
		N = times
	}
	dp := make([][]float64, N+1)
	for i:=0; i<N+1; i++ {
		dp[i] = make([]float64, N+1)
	}
	dp[0][0] = 0.5
	for i:=1; i<N+1; i++ {
		dp[0][i] = 1.0
		dp[i][0] = 0.0
	}
	for i:=1; i<N+1; i++ {
		for j:=1; j<N+1; j++ {
			sum := 0.0
			if i-4 < 0 {
				sum += dp[0][j]
			} else {
				sum += dp[i-4][j]
			}
			if i-3 < 0 {
				if j-1 < 0 {
					sum += dp[0][0]
				} else {
					sum += dp[0][j-1]
				}
			} else {
				if j-1 < 0 {
					sum += dp[i-3][0]
				} else {
					sum += dp[i-3][j-1]
				}
			}
			if i-2 < 0 {
				if j-2 < 0 {
					sum += dp[0][0]
				} else {
					sum += dp[0][j-2]
				}
			} else {
				if j-2 < 0 {
					sum += dp[i-2][0]
				} else {
					sum += dp[i-2][j-2]
				}
			}
			if i-1 < 0 {
				if j-3 < 0 {
					sum += dp[0][0]
				} else {
					sum += dp[0][j-3]
				}
			} else {
				if j-3 < 0 {
					sum += dp[i-1][0]
				} else {
					sum += dp[i-1][j-3]
				}
			}
			dp[i][j] =0.25 * sum
		}
	}
	return dp[N][N]
}
