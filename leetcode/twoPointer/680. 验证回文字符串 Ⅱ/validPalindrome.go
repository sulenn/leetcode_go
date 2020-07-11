package main

//https://leetcode-cn.com/problems/valid-palindrome-ii/

//这个题用双指针，当遇到s[i]!=s[j]时不能通过s[i+1]==s[j]或s[i]==s[j-1]来判断删除哪一个字符，
//而是需要分两种情况来讨论（即删除左边或删除右边）
func validPalindrome(s string) bool {
	p1 := 0
	p2 := len(s) - 1
	flag := 1                // 可删除一次
	record := make([]int, 2) // 用于记录删除右边
	for p1 <= p2 {
		if s[p1] == s[p2] {
			p1++
			p2--
		} else if flag > 0 { // 删除左边
			record[0] = p1
			record[1] = p2 - 1
			p1++
			flag--
		} else {
			break
		}
	}
	if p1 > p2 {
		return true
	} //不用删除字符、或删除左边字符回文成功
	p1 = record[0]
	p2 = record[1]
	for p1 <= p2 { // 删除右边字符
		if s[p1] == s[p2] {
			p1++
			p2--
		} else {
			return false
		}
	}
	return true
}
