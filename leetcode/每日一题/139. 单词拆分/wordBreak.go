package main

//https://leetcode-cn.com/problems/word-break/

// 回溯，超时
//func wordBreak(s string, wordDict []string) bool {
//	wordDicts := make(map[string]struct{})
//	for i:=0; i<len(wordDict); i++ {
//		wordDicts[wordDict[i]] = struct{}{}
//	}
//	return backtracking(s, wordDicts)
//}
//
//func backtracking(s string, wordDicts map[string]struct{}) bool {
//	if len(s) == 0 {
//		return true
//	}
//	for i:=0; i<len(s); i++ {
//		word := s[:i+1]
//		if _, ok := wordDicts[word]; ok {
//			if backtracking(s[i+1:], wordDicts) {
//				return true
//			}
//		}
//	}
//	return false
//}

//dp:
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	word_dict := make(map[string]bool)
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