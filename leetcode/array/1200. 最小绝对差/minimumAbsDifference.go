package main

import (
	"math"
	"sort"
)

//https://leetcode-cn.com/problems/minimum-absolute-difference/

//排序之后，再处理
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr) // 先排序
	result := [][]int{}
	min := math.MaxInt64
	for i := 0; i < len(arr)-1; i++ {
		value := arr[i] - arr[i+1]
		if value < 0 {
			value = -value
		}
		if value < min {
			min = value
			result = [][]int{{arr[i], arr[i+1]}}
		} else if value == min {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}
