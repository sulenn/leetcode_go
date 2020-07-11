package main

import "fmt"

//https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

func firstUniqChar(s string) byte {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; ok {
			sMap[s[i]]++
		} else {
			sMap[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	test := make(map[int]int)
	fmt.Println(test[0])
}
