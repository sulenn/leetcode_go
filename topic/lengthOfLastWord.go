package main

import (
	"fmt"
)

//func lengthOfLastWord(s string) int {
//	wordArray := strings.Split(s, " ")
//	for i := len(wordArray) - 1; i >= 0; i-- {
//		if len(wordArray[i]) != 0{
//			fmt.Println(wordArray[i])
//			return len(wordArray[i])
//		}
//	}
//	return 0
//}

func lengthOfLastWord(s string) int {
	length := 0    //记录最后一个单词的长度
	for i := len(s) - 1;i >= 0;i-- {
		if length != 0 && s[i] == ' ' {
			return length
		} else if s[i] != ' ' {
			length++
		}
	}
	return length
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World  "))
}
