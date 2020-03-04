package main

import (
	"strconv"
)

//https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/

//通过在排序时传入一个自定义的 Comparator 实现，重新定义 String 列表内的排序方法，
//若拼接 s1 + s2 > s2 + s1，那么显然应该把 s2 在拼接时放在前面，以此类推，将整个 String 列表排序后再拼接起来。
func minNumber(nums []int) string {
	for i:=0; i<len(nums); i++ {  // O（n）的排序方式。可以用堆排序或者快排
		for j:=i+1; j<len(nums); j++ {
			if strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) >= strconv.Itoa(nums[j])+strconv.Itoa(nums[i]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	result := ""
	for i:=0 ;i<len(nums); i++ {
		result += strconv.Itoa(nums[i])
	}
	return result
}

