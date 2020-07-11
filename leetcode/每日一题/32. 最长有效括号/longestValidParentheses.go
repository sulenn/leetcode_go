package main

//https://leetcode-cn.com/problems/longest-valid-parentheses/

//参考：https://leetcode-cn.com/problems/longest-valid-parentheses/solution/shou-hua-tu-jie-zhan-de-xiang-xi-si-lu-by-hyj8/

func longestValidParentheses(s string) int {
	stack := []int{-1} // 初始值
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		if s[i] == ')' && len(stack) > 1 {
			if i-stack[len(stack)-2] > max {
				max = i - stack[len(stack)-2]
			}
			stack = stack[:len(stack)-1]
		} else if s[i] == ')' {
			stack = []int{i}
		}
	}
	return max
}
