package main

import "sort"

//https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/

// 全排列， 使用 hashtable 去重
//func permutation(s string) []string {
//	if len(s) == 1 {return []string {s}}
//	result := []string {}
//	table := make(map[byte]interface{})   // 使用 hashtable 来去重
//	for i:=0; i < len(s); i++ {  // 递归
//		if _, ok := table[s[i]]; !ok {
//			strArr := permutation(s[:i] + s[i+1:])
//			for j:=0; j<len(strArr); j++ {
//				strArr[j] = string(s[i]) + strArr[j]
//			}
//			result = append(result, strArr...)
//			table[s[i]] = nil
//		}
//	}
//	return result
//}

// 参考网友
func permutation(s string) []string {
	if len(s) <= 0 {
		return []string{}
	}
	n := len(s)
	pickedList := make([]bool, n)
	cur := make([]byte, n)
	res := []string{}
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	var permuteHelper func(int)
	permuteHelper = func(index int) {
		if index >= n {
			res = append(res, string(cur))
		}
		for i := 0; i < n; i++ {
			if !pickedList[i] && (i == 0 || bs[i-1] != bs[i] || !pickedList[i-1]) {
				pickedList[i] = true
				cur[index] = bs[i]
				permuteHelper(index + 1)
				pickedList[i] = false
			}
		}
	}
	permuteHelper(0)
	return res
}
