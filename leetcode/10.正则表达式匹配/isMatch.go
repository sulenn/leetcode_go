package main

import "fmt"

//https://leetcode-cn.com/problems/regular-expression-matching/

// 参考：《剑指offer》
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(p) == 0 {
		return false
	}
	if p[0] == '*' {
		return isMatch(s, p[1:])
	}
	if len(p) > 1 && p[1] == '*' {
		if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
			return isMatch(s, p[2:]) || isMatch(s[1:], p) || isMatch(s[1:], p[2:])
		}
		return isMatch(s, p[2:])
	}
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		return isMatch(s[1:], p[1:])
	}
	return false
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("a", "ab*"))
	fmt.Println(isMatch("ab", ".*c"))
}
