package main

import "math"

//https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

func maxSubArray(nums []int) int {
	result := math.MinInt64
	sum := 0
	for _, value := range nums {
		if sum + value >= 0 {
			sum += value
			if sum > result {
				result = sum
			}
		} else {
			sum = 0
			if value > result {
				result = value
			}
		}
	}
	return result
}
