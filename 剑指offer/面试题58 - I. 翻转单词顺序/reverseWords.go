package main

import "strings"

//https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	result := ""
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] != "" {
			result += sArr[i] + " "
		}
	}
	return strings.Trim(result, " ")
}
