package main

//https://leetcode-cn.com/problems/reverse-words-in-a-string/

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	newArr := make([]string, 0)
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] != "" {
			newArr = append(newArr, sArr[i])
		}
	}
	return strings.Join(newArr, " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords("  a good   example  "))
}
