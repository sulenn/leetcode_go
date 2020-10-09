package main

//https://leetcode-cn.com/problems/new-21-game/

//回溯暴力
//var sumPro float64
//func new21Game(N int, K int, W int) float64 {
//	sumPro = 0
//	backtracking(N, K, W, 0,1.0)
//	return sumPro
//}
//
//func backtracking(N int, K int, W int, sum int, pro float64)  {
//	if sum >= K {
//		if sum <= N {
//			sumPro += pro
//		}
//		return
//	}
//	for i:=1; i<=W; i++ {
//		backtracking(N, K, W, sum+i, pro*(1.0/float64(W)))
//	}
//}

// 逆向动态规划：https://leetcode-cn.com/problems/new-21-game/solution/huan-you-bi-zhe-geng-jian-dan-de-ti-jie-ma-tian-ge/
func new21Game(N int, K int, W int) float64 {
	if K > N {
		return 0
	}
	dp := make([]float64, K+W)
	proSum := 0.0 // 注意和
	for i := K; i < K+W; i++ {
		if i <= N {
			dp[i] = 1
		} else {
			dp[i] = 0
		}
		proSum += dp[i]
	}

	for i := K - 1; i >= 0; i-- {
		dp[i] = proSum / float64(W)
		proSum = proSum - dp[i+W] + dp[i]
	}
	return dp[0]
}
