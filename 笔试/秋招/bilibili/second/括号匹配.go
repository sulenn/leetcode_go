package main

import "strings"

func IsValidExp(s string) bool {
	// write code here
	stack := make([]byte, 0)
	left := ")]}"
	for i := 0; i < len(s); i++ {
		if strings.Contains(left, string(s[i])) {
			if len(stack) == 0 {
				return false
			}
			if s[i] == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				continue
			}
			if s[i] == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				continue
			}
			if s[i] == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
