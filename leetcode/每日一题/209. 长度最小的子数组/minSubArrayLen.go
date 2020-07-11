package main

import "math"

//https://leetcode-cn.com/problems/minimum-size-subarray-sum/

// 双指针
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	rel := math.MaxInt64
	first := 0
	sum := 0
	for i := 0; i <= len(nums); i++ {
		for sum >= s {
			if i-first < rel {
				rel = i - first
			}
			if first == i-1 {
				return rel
			}
			sum -= nums[first]
			first++
		}
		if i == len(nums) { // 数组中最后一个数，边界值溢出
			break
		}
		sum += nums[i]
	}
	if rel == math.MaxInt64 {
		return 0
	}
	return rel
}
