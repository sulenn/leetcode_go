package main

import "strconv"

//https://leetcode-cn.com/problems/decode-string/

//最简洁的方法，参考：https://leetcode-cn.com/problems/decode-string/solution/decode-string-fu-zhu-zhan-fa-di-gui-fa-by-jyd/
// 以下实现是自己的写的垃圾，没有参考链接
func decodeString(s string) string {
	numArr := make([]int, 0)
	bracketNum := 0
	alphaArr := make([][]byte, 1)  // 注意从1开始，假如不存在括号
	for i:=0; i<len(s); i++{
		if s[i] >= '0' && s[i] <= '9' {  // 数字可能不止一位
			start := i
			i++
			for s[i] >= '0' && s[i] <= '9' {
				i++
			}
			num, _ := strconv.Atoi(s[start:i])
			numArr = append(numArr, num)
		}
		if s[i] == '[' {   // 左括号
			bracketNum++
			alphaArr = append(alphaArr, []byte {})
		} else if s[i] == ']' {   // 右括号
			num := numArr[len(numArr)-1]
			numArr = numArr[:len(numArr)-1]
			alpha := alphaArr[len(alphaArr)-1]
			alphaArr = alphaArr[:len(alphaArr)-1]
			for j:=0; j<num; j++ {
				alphaArr[len(alphaArr)-1] = append(alphaArr[len(alphaArr)-1], alpha...)
			}
			bracketNum--
		} else {
				alphaArr[len(alphaArr)-1] = append(alphaArr[len(alphaArr)-1], s[i])
		}
	}
	return string(alphaArr[0])
}
