package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strArray := strings.Split(s, " ")
	result := ""
	for i := 0; i < len(strArray); i++ {
		temp := ""
		for j := 0; j < len(strArray[i]); j++ {
			temp = string(strArray[i][j]) + temp
		}
		result += temp + " "
	}
	return result[:len(result)-1]
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
