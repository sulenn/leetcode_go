package main

import "fmt"

//https://leetcode-cn.com/problems/valid-number/

// 参考《剑指 offer》

func isNumber(s string) bool {
	for len(s) != 0 && s[0] == ' ' { // 过滤前导空格
		s = s[1:]
	}
	for len(s) != 0 && s[len(s)-1] == ' ' { // 过滤后导空格
		s = s[:len(s)-1]
	}
	index := 0
	flag := judgeInteger(s, &index)
	if index < len(s) && s[index] == '.' {
		index++
		flag = judgeUnsignedInteger(s, &index) || flag
	}
	if index < len(s) && (s[index] == 'e' || s[index] == 'E') {
		index++
		flag = judgeInteger(s, &index) && flag
	}
	return flag && index == len(s)
}

func judgeInteger(s string, index *int) bool {
	if *index < len(s) && (s[*index] == '+' || s[*index] == '-') {
		*index++
	}
	return judgeUnsignedInteger(s, index)
}

func judgeUnsignedInteger(s string, index *int) bool {
	initIndex := *index
	for *index < len(s) && s[*index] >= '0' && s[*index] <= '9' {
		*index++
	}
	return *index > initIndex
}

func main() {
	fmt.Println(isNumber("0"))
	fmt.Println(isNumber(" 0.1"))
	fmt.Println(isNumber("abc"))
	fmt.Println(isNumber("1 a"))
	fmt.Println(isNumber("2e10"))
	fmt.Println(isNumber(" -90e3"))
	fmt.Println(isNumber(" 1e"))
	fmt.Println(isNumber("e3"))
	fmt.Println(isNumber(" 6e-1"))
	fmt.Println(isNumber(" 99e2.5 "))
	fmt.Println(isNumber("53.5e93"))
	fmt.Println(isNumber(" --6"))
	fmt.Println(isNumber("-+3"))
	fmt.Println(isNumber("95a54e53"))
	fmt.Println(isNumber("1 "))
	fmt.Println(isNumber(".1"))
	fmt.Println(isNumber("3."))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber(""))
}
