package main

//https://leetcode-cn.com/problems/longest-palindromic-subsequence/

func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {return 0}
	length := len(s)
	arr := make([][]int, length)
	for i:=0; i<length; i++ {   // 初始化
		arr[i] = make([]int, length)
	}
	for i:=length-1; i>=0; i-- {
		for j:=i; j<length; j++ {
			if i==j {  //对角线为1
				arr[i][j] = 1
				continue
			}
			if j - i == 1 && s[i] == s[j] {   // i 和 j 相邻
				arr[i][j] = 2
				continue
			}
			if s[i] == s[j] {   // dp
				arr[i][j] = arr[i+1][j-1] + 2
			} else {
				arr[i][j] = max(arr[i+1][j], arr[i][j-1])
			}
		}
	}
	return arr[0][length-1]
}

func max(a, b int) int {
	if a > b {return a}
	return b
}