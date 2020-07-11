package main

//https://leetcode-cn.com/problems/longest-palindrome/

func longestPalindrome(s string) int {
	alpha := make([]int, 52)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' { // 大写范围 0-25
			alpha[s[i]-'A']++
		} else { // 小写范围 26 - 51
			alpha[s[i]-'a'+26]++
		}
	}
	result := 0
	for _, v := range alpha {
		result += v / 2 * 2
	}
	if result != len(s) { // 存在奇数
		return result + 1
	} else { // 不存在奇数
		return result
	}
}
