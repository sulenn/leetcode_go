package main

//https://leetcode-cn.com/problems/valid-parenthesis-string/

func checkValidString(s string) bool {
	sums := 0 //  已匹配的左括号和星号数量
	left := 0 // 剩余未被匹配的左括号数量
	for _, v := range s {
		if v == '(' {
			left++
		} else if v == '*' { // 星号
			if left != 0 {
				left--
				sums++
			}
			sums++
		} else if v == ')' { //右括号
			if left != 0 { // 左括号还有剩余
				left--
			} else if sums != 0 { // 星号和左括号有剩余
				sums--
			} else {
				return false
			}
		}
	}
	if left != 0 { // 还有左括号
		return false
	}
	return true
}
