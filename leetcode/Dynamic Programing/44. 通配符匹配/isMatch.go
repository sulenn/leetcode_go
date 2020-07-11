package main

//https://leetcode-cn.com/problems/wildcard-matching/

//可以贪心和动态规划两种解法

//参考：
//https://leetcode-cn.com/problems/wildcard-matching/solution/javashuang-zhi-zhen-tan-xin-he-dong-tai-gui-hua-xi/
//https://leetcode-cn.com/problems/wildcard-matching/solution/44-tong-pei-fu-pi-pei-shuang-zhi-zhen-by-guohaodin/

//dp，思路参考链接1
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return p == s
	} // 模式串为空，如果模式串和字符串均为空则 true
	dp := make([][]bool, len(s)+1) // 初始化 dp 数组
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true                 // s前0个字符和p前0个字符必定匹配
	for i := 1; i < len(p)+1; i++ { // dp[0][1]~dp[0][p.length]只有p的i字符以及前面所有字符都为'*'才为true
		dp[0][i] = p[i-1] == '*' && dp[0][i-1]
	}
	for i := 1; i < len(s)+1; i++ { // 遍历
		for j := 1; j < len(p)+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
