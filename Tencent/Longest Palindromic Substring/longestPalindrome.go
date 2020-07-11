package main

//https://leetcode-cn.com/problems/longest-palindromic-substring/submissions/

import "fmt"

//动态规划的方法
func longestPalindrome(s string) string {
	palindrome := ""
	length := 0
	//单节点
	for i := 0; i < len(s); i++ {
		result := judgementSingle(i, s)
		if result*2+1 > length {
			palindrome = s[i-result : i+result+1] // 注意右边界取值
			length = result*2 + 1
		}
	}
	//双节点
	for j := 0; j < len(s)-1; j++ {
		if s[j] == s[j+1] { // 先判断两个值是否相等
			result := judgementDouble(j, j+1, s)
			if result*2+2 > length {
				palindrome = s[j-result : j+1+result+1]
				length = result*2 + 2
			}
		}
	}
	return palindrome
}

//以单个点为中心
func judgementSingle(index int, s string) int {
	count := 0
	for index-1-count >= 0 && index+1+count < len(s) {
		if s[index-1-count] == s[index+1+count] {
			count++
		} else {
			break
		}
	}
	return count
}

//以两个点为中心
func judgementDouble(pre int, back int, s string) int {
	count := 0
	for pre-1-count >= 0 && back+1+count < len(s) {
		if s[pre-1-count] == s[back+1+count] {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
