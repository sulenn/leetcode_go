package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

//用字典记录每个字符出现的位置，滑动窗口
func lengthOfLongestSubstring(s string)int {
	result,maxRepeatIndex := 0,0
	var m = map[string]int{}
	for i := 0; len(s) > i; i++ {
		if _, ok := m[string(s[i])]; ok {
			// 首先记住如果出现重复距离重复最近的 位置
			maxRepeatIndex = max(maxRepeatIndex,m[string(s[i])])
		}
		// 更新每次重复数字的最大值
		result = max(i - maxRepeatIndex + 1 , result)
		m[string(s[i])] = i + 1
	}
	fmt.Println(result)
	return result
}
func max(a,b int)int{
	if a < b {
		return b
	}
	return a
}

// 字符串滑动窗口解法，不够优雅。时间复杂度O(n*2),空间复杂度O(n)
//func lengthOfLongestSubstring(s string) int {
//	result := 0
//	longStr := ""
//	length := len(s)
//	for i:=0; i<length; i++ {
//		index := strings.Index(longStr, string(s[i]))
//		if index == -1 {
//			longStr += string(s[i])
//		} else {
//			if len(longStr) > result {
//				result = len(longStr)
//			}
//			longStr = longStr[index+1:] + string(s[i])
//		}
//	}
//	if len(longStr) > result {
//		result = len(longStr)
//	}
//	return result
//}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	//fmt.Println(strings.Index("widuu", "i")) //1
	//fmt.Println(strings.Index("widuu", "a")) //3
}
