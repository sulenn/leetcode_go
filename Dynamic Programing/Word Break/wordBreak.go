package main

import "fmt"

//https://leetcode-cn.com/problems/word-break/

//参考动态规划：https://leetcode-cn.com/problems/word-break/solution/dong-tai-gui-hua-ji-yi-hua-hui-su-zhu-xing-jie-shi/

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	word_dict := make(map[string]bool)    // 字典
	for _, value := range wordDict {
		word_dict[value] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i:=0; i<len(s); i++ {
		if dp[i] != true {
			continue
		}
		for j:=i+1; j<len(s)+1; j++ {
			if _, ok := word_dict[s[i:j]]; ok {
				dp[j] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string {"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string {"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string {"cats", "dog", "sand", "and", "cat"}))
}
