package main

import "fmt"

//https://leetcode-cn.com/problems/valid-number/

func isNumber(s string) bool {
	for len(s) != 0 && s[0] == ' ' { // 过滤前导空格
		s = s[1:]
	}
	for len(s) != 0 && s[len(s)-1] == ' ' { // 过滤后导空格
		s = s[:len(s)-1]
	}
	if len(s) == 0 {
		return false
	}
	length := len(s)
	index := 0
	if !scanA(s, length, &index) { // 满足 A
		return false
	}
	if index == length {
		return true
	}
	if length == 1 && s == "." {
		return false
	}
	if s[index] == '.' { // 满足 B
		index++
		if !scanB(s, length, &index) {
			return false
		}
	}
	if index == length {
		return true
	}

	if index != 0 && (s[index] == 'e' || s[index] == 'E') { // 满足 C
		index++
		if !scanC(s, length, &index) {
			return false
		}
	}
	return index == length
}

func scanA(s string, length int, index *int) bool {
	for *index < length {
		if (s[*index] == '+' || s[*index] == '-') && *index != 0 {
			return false
		}
		if s[*index] == '.' {
			return true
		}
		if s[*index] == 'e' || s[*index] == 'E' {
			return true
		}
		if s[*index] < '0' || s[*index] > '9' {
			return false
		}
		*index++
	}
	return true
}

func scanB(s string, length int, index *int) bool {
	for *index < length {
		if s[*index] == 'e' || s[*index] == 'E' {
			return true
		}
		if s[*index] < '0' || s[*index] > '9' {
			return false
		}
		*index++
	}
	return true
}

func scanC(s string, length int, index *int) bool {
	if length == *index {
		return false
	}
	initIndex := index
	for *index < length {
		if (s[*index] == '+' || s[*index] == '-') && index == initIndex {
			*index++
			continue
		}
		if s[*index] < '0' || s[*index] > '9' {
			return false
		}
		*index++
	}
	return true
}

func main() {
	//fmt.Println(isNumber("0"))
	//fmt.Println(isNumber(" 0.1"))
	//fmt.Println(isNumber("abc"))
	//fmt.Println(isNumber("1 a"))
	//fmt.Println(isNumber("2e10"))
	//fmt.Println(isNumber(" -90e3"))
	//fmt.Println(isNumber(" 1e"))
	//fmt.Println(isNumber("e3"))
	//fmt.Println(isNumber(" 6e-1"))
	//fmt.Println(isNumber(" 99e2.5 "))
	//fmt.Println(isNumber("53.5e93"))
	//fmt.Println(isNumber(" --6"))
	//fmt.Println(isNumber("-+3"))
	//fmt.Println(isNumber("95a54e53"))
	fmt.Println(isNumber("1 "))
	fmt.Println(isNumber(".1"))
	fmt.Println(isNumber("3."))
	fmt.Println(isNumber("."))
}
