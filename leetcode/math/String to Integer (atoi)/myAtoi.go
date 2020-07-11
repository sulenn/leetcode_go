package main

import "fmt"

//https://leetcode-cn.com/problems/string-to-integer-atoi/

//这个题目好蠢
func myAtoi(str string) int {
	for len(str) != 0 && str[0] == ' ' { // 防止溢出
		str = str[1:]
	}
	if len(str) == 0 { // 字符串全由空格组成
		return 0
	}
	if str[0] != '+' && str[0] != '-' && (str[0] > '9' || str[0] < '0') { // +1，输出为1
		return 0
	}
	flag := false      // 正负号
	if str[0] == '-' { // 负号
		flag = true
		str = str[1:]
	} else if str[0] == '+' { // 两个符号叠在一起返回0
		str = str[1:]
	}
	result := 0
	for i := 0; i < len(str); i++ {
		if str[i] <= '9' && str[i] >= '0' {
			result = result*10 + int(str[i]-'0')
		} else { // 3.142
			break
		}
		if result > 2147483648 { // 范围是-2147483648 ~ 2147483647
			result = 2147483648
			break
		}
	}
	if flag {
		result = -result
	}
	if result > 2147483647 { //正数比负数少1
		return 2147483647
	}
	return result
}

func main() {
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("3.14159"))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("-"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("-+1"))
	fmt.Println(myAtoi("2147483648"))
}
