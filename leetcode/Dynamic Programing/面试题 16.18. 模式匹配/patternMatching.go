package main

//https://leetcode-cn.com/problems/pattern-matching-lcci/

func patternMatching(pattern string, value string) bool {
	for i:=0; i<=len(value); i++ {
		for j:=i+1; j<=len(value); j++ {
			if judgement(value[:i], value[i:j], pattern, value) {
				return true
			}
		}
	}
	return false
}

func judgement(a string, b string, pattern string, value string) bool {
	j := 0   // value 索引下标
	lenA := len(a)
	lenB := len(b)
	lenValue := len(value)
	for i:=0; i<len(pattern); i++ {
		if pattern[i] == 'a' && j+lenA <= lenValue && a == value[j:j+lenA] {
			j += lenA
			continue
		}
		if pattern[i] == 'b' && j+lenB <= lenValue &&  b == value[j:j+lenB] {
			j += lenB
			continue
		}
		return false
	}
	if j != lenValue {
		return false
	}
	return true
}

//func backtracking()  {
//
//}