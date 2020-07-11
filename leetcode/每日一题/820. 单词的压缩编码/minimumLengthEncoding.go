package main

import (
	"sort"
	"strings"
)

//https://leetcode-cn.com/problems/short-encoding-of-words/

//思路，将所有单词反转，然后排序。原单词中后缀相同的单词进行匹配即可
func minimumLengthEncoding(words []string) int {
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}
	sort.Strings(words) // 反转后的字符串排序
	result := 0
	for i := 0; i < len(words)-1; i++ {
		if strings.Contains(words[i+1], words[i]) && words[i] == words[i+1][:len(words[i])] { // 两个单词的后缀相同
			continue
		} else {
			result += len(words[i]) + 1
		}
	}
	result += len(words[len(words)-1]) + 1 // 最后一个单词
	return result
}

func reverse(str string) string { // 反转字符串
	arr := []rune(str)
	l, h := 0, len(str)-1
	for l < h {
		arr[l], arr[h] = arr[h], arr[l]
		l++
		h--
	}
	return string(arr)
}
