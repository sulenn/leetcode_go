package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)   //每次递归返回的字符串
	newStr := ""
	numCount := 1
	currentNum := str[0]
	for i := 1;i < len(str); i++ {
		if currentNum == str[i] {
			numCount++
		} else {
				newStr = newStr + strconv.Itoa(numCount) + string(currentNum)
				currentNum = str[i]
				numCount = 1
		}
	}
	newStr = newStr + strconv.Itoa(numCount) + string(currentNum)
	return newStr
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(4))
}
