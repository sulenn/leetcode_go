package main

//https://leetcode-cn.com/problems/unique-binary-search-trees/

// dp[i] = dp[i-1]*dp[0] + dp[i-2]*dp[1] + dp[i-3]*dp[2] + ... + dp[1]*dp[i-2] + dp[0]*dp[i-1]
//可参考：https://leetcode-cn.com/problems/unique-binary-search-trees/solution/hua-jie-suan-fa-96-bu-tong-de-er-cha-sou-suo-shu-b/
func numTrees(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	for i:=1; i<=n; i++ {
		for j:=1; j<=i/2;j++ {   // 状态方程中的每一项满足对称，偶数
			arr[i] += 2*arr[i-j]*arr[j-1]
		}
		if i%2 != 0 {  // 奇数
			arr[i] += arr[i/2]*arr[i/2]
		}
	}
	return arr[n]
}
