package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/

//题中“字典顺序最小的字符串”指的是按字符比较大小时最小，而不是字符串在初始字典中的顺序小
//比如 "bab"，["ba","ab","a","b"] 输入的预期输出为 "ab"
type str struct {
	s string
	l int
}

//思路：先将字符串数组按长度从大到小排列，然后同等长度的字符串再按字符大小排列
//然后开始和模式串匹配
func findLongestWord(s string, d []string) string {
	strArr := make([]str, len(d))
	for k, v := range d {
		strArr[k] = str{v, len(v)}
	}
	sort.Slice(strArr, func(i, j int) bool {  // 字符串按长度排列
		return strArr[i].l > strArr[j].l
	})
	p1 := 0
	for i:=0; i<len(strArr)-1; i++ {   // 在长度重大到小的基础上，同等长度的字符串按字符大小排列
		if len(strArr[i+1].s ) != len(strArr[i].s) {
			if i > p1 {
				sortByWord(strArr[p1:i+1])   // 同等长度的字符串按大小排列
			}
			p1 = i+1
		}
	}
	sortByWord(strArr[p1:])   // 同等长度的字符串按大小排列
	for i:=0; i<len(strArr); i++ {  // 字符串匹配
		ps := 0  // 模式串指针
		pd := 0  // 另一个指针
		for ;ps <len(s) && pd < len(strArr[i].s); ps++ {
			if strArr[i].s[pd] == s[ps] {
				pd++
			}
		}
		if pd == len(strArr[i].s) {return strArr[i].s}
	}
	return ""
}

func sortByWord(s []str)  {  // 字符串按字符大小排列
	sort.Slice(s, func(i, j int) bool {
		return s[i].s < s[j].s
	})
}