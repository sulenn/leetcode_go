package main

import "strconv"

//https://leetcode-cn.com/problems/compress-string-lcci/

func compressString(S string) string {
	if len(S) < 3 {return S}
	result := ""
	cur := S[0]
	num := 1
	for i:=1; i<len(S); i++ {
		if cur == S[i] {
			num++
		} else {
				result += string(cur) + strconv.Itoa(num)
				cur = S[i]
				num = 1
		}
	}
	result += string(cur) + strconv.Itoa(num)
	if len(result) >= len(S) {
		return S
	}
	return result
}
