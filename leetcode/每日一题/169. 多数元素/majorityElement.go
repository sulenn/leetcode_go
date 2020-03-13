package main

import "math"

//https://leetcode-cn.com/problems/majority-element/

func majorityElement(nums []int) int {
	result := math.MaxInt64
	sum := 0
	for _, v := range nums {
		if result == v {sum++}
		if result != v && sum > 0 {sum--}
		if result != v && sum <= 0 {
			result = v
			sum = 1
		}
	}
	return result
}
