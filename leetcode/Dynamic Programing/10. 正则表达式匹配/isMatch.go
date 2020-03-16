package main

//https://leetcode-cn.com/problems/regular-expression-matching/

//动态规划，参考：https://leetcode-cn.com/problems/regular-expression-matching/solution/dong-tai-gui-hua-zen-yao-cong-0kai-shi-si-kao-da-b/

//这道题太难了。思路和通配符匹配差不多，但是递归的边界处理更复杂
func isMatch(s string, p string) bool {
	if len(p) == 0 {return p == s}
	dp := make([][]bool, len(s)+1)  // 初始化 dp 数组
	for i:=0; i<len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true   // s前0个字符和p前0个字符必定匹配
	for i:=1; i<len(p)+1; i++ {
		if p[i-1] == '*' && dp[0][i-1] {
			dp[0][i] = dp[0][i-1]
		} else if p[i-1] == '*' && !dp[0][i-1]  {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i:=1; i<len(s)+1; i++ {
		for j:=1; j<len(p)+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' { // s[i] = p[j] 或者 p[j] = '.'
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' { // 分类讨论
				if j-2 < 0 { // 防止溢出
					dp[i][j] = dp[i][j-1]
				} else if p[j-2] == s[i-1] || p[j-2] == '.' { // * 前一个字符和 s[i] 相同，或者 p[j-1] = ‘.'
					dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
				} else if p[j-2] != s[i-1] { // * 前一个字符和 s[i] 不同
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
