package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/contest/weekly-contest-172/problems/print-words-vertically/

func printVertically(s string) []string {
	sArray := strings.Split(s, " ")   // 字符串切片
	sLength := make([]int, len(sArray))    // 每个小字符串的长度，避免溢出
	maxLength := 0   // 最长的小字符串长度
	for i:=0; i < len(sArray); i++ {
		sLength[i] = len(sArray[i])
		if maxLength < len(sArray[i]) {
			maxLength = len(sArray[i])
		}
	}
	result := make([]string, maxLength)
	for j:=0; j<maxLength; j++ {
		temp_str := ""
		for index, str := range sArray {   // 拼接竖立字符串
			if j < sLength[index] {
				temp_str += string(str[j])
			} else {
				temp_str += " "
			}
		}
		spaceNum := 0
		for i := len(temp_str)-1; i >= 0; i-- {  // 去除字符串尾部的所有空格
			if temp_str[i] == ' ' {
				spaceNum++
			} else {
					break
			}
		}
		result[j] = temp_str[:len(temp_str)-spaceNum]
	}
	return result
}

func main() {
	fmt.Println(printVertically("HOW ARE YOU"))
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
	fmt.Println(printVertically("CONTEST IS COMING"))
}
