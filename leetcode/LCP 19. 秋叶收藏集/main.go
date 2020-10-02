package main

//https://leetcode-cn.com/problems/UlBDOe/

//参考：https://leetcode-cn.com/problems/UlBDOe/solution/dong-tai-gui-hua-java-qiu-xie-shou-cang-ji-by-crus/
//https://leetcode-cn.com/problems/UlBDOe/solution/ulbdoe-by-ikaruga/
func minimumOperations(leaves string) int {
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, len(leaves))
	}
	for i := 0; i < len(leaves); i++ { // 第一行，全红
		if i == 0 {
			dp[0][i] = isYellow(leaves[i])
			continue
		}
		dp[0][i] = dp[0][i-1] + isYellow(leaves[i])
	}
	for i := 0; i < len(leaves); i++ { // 第二行，红黄
		if i == 0 {
			dp[1][i] = dp[0][i]
			continue
		}
		dp[1][i] = min(dp[0][i-1], dp[1][i-1]) + isRed(leaves[i])
	}
	for i := 0; i < len(leaves); i++ { // 第三行，红黄红
		if i == 0 || i == 1 {
			dp[2][i] = dp[1][i]
			continue
		}
		dp[2][i] = min(dp[1][i-1], dp[2][i-1]) + isYellow(leaves[i])
	}
	return dp[2][len(leaves)-1]
}

func isYellow(b byte) int {
	if b == 'y' {
		return 1
	}
	return 0
}

func isRed(b byte) int {
	if b == 'r' {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
