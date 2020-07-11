package main

//https://leetcode-cn.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}
	min := ""
	alphaMap := make(map[byte]int)
	alphaSumInT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		alphaSumInT[t[i]]++
	}
	valid := 0
	left, right := 0, 0
	for ; right < len(s); right++ {
		if sum, ok := alphaSumInT[s[right]]; ok {
			alphaMap[s[right]]++
			if alphaMap[s[right]] == sum {
				valid++
			}
		}
		for valid == len(alphaSumInT) {
			if len(min) == 0 {
				min = s[left : right+1]
			} else if len(min) > right-left+1 {
				min = s[left : right+1]
			}
			if sum, ok := alphaSumInT[s[left]]; ok {
				if alphaMap[s[left]] == sum {
					valid--
				}
				alphaMap[s[left]]--
			}
			left++
		}
	}
	return min
}
