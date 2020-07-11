package main

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/multiply-strings/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" { // 特殊情况
		return "0"
	}
	result := ""
	j := 0 // 第几位，代表要多加几个零
	for i := len(num1) - 1; i >= 0; i-- {
		singleMulti := singleMutiply(num1[i], num2, j) // num1 的每一位和 num2 相乘
		result = add(result, singleMulti)              // 乘完相加
		j++
	}
	return result
}

//单个数字和字符串相乘
func singleMutiply(num1 uint8, num2 string, num int) string {
	carry := 0
	result := ""
	for i := len(num2) - 1; i >= 0; i-- {
		multi := int((num1 - '0') * (num2[i] - '0'))
		rest := multi % 10
		result = strconv.Itoa((rest+carry)%10) + result // 注意 (rest+carry)%10
		carry = (rest+carry)/10 + multi/10
	}
	if carry != 0 {
		result = strconv.Itoa(carry) + result
	}
	for j := 0; j < num; j++ {
		result += "0"
	}
	return result
}

//两个字符串相加
func add(num1 string, num2 string) string {
	if len(num1) < len(num2) { // 取最短
		num1, num2 = num2, num1
	}
	for len(num1) != len(num2) { // 用零补齐
		num2 = "0" + num2
	}
	carry := 0
	result := ""
	for i := len(num1) - 1; i >= 0; i-- {
		sum := int((num1[i]-'0')+(num2[i]-'0')) + carry // uint8 类型减去 '0'
		carry = sum / 10
		rest := sum % 10
		result = strconv.Itoa(rest) + result
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}

func main() {
	//fmt.Println(add("123", "456"))
	//fmt.Println(add("0", "0"))
	//fmt.Println(singleMutiply('8', "456", 1))
	//fmt.Println(multiply("2", "3"))
	//fmt.Println(multiply("123", "456"))
	//fmt.Println(multiply("3", "456"))
	//fmt.Println(multiply("0", "456"))
	//fmt.Println(multiply("9", "987654321"))
	//fmt.Println(multiply("8", "987654321"))
	fmt.Println(multiply("123456789", "987654321"))
}
