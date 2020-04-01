package main

//https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/

func maxDepthAfterSplit(seq string) []int {
	var stack int
	result := make([]int, len(seq))
	for i:=0; i<len(seq); i++ {
		if seq[i] == '(' {
			stack++
		}
		if stack % 2 == 1 {    // 括号栈深度为奇数的属于A
			result[i] = 0
		} else {   // 括号栈深度为偶数的属于B
			result[i] = 1
		}
		if seq[i] == ')' {
			stack--
		}
	}
	return result
}
