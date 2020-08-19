package main

import "fmt"

func countSubstrings(s string) int {
	length := len(s)
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}
	result := length
	for i := length - 2; i >= 0; i-- { // 注意开始的位置，倒序
		for j := i + 1; j < length; j++ {
			if j-i == 1 && s[i] == s[j] { // 偶数，且 j-i =1，也就是两个字符相邻
				result++
				dp[i][j] = true
				continue
			}
			if s[i] == s[j] && dp[i+1][j-1] {
				result++
				dp[i][j] = true
				continue
			}
		}
	}
	return result
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings(""))
}
