package main

import "math"

//https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/

//满足要求的 number 数量超过一半，参考剑指 offer
func majorityElement(nums []int) int {
	count := 0
	number := math.MaxInt64
	for _, value := range nums {
		if number != value && count > 1 {
			count--
		} else if number != value {
			count = 1
			number = value
		} else {
			count++
		}
	}
	return number
}
