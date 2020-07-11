package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/3sum-closest/

// 思路：排序数组。固定一个值，然后双指针遍历
func threeSumClosest(nums []int, target int) int {
	sum := 99999
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > target {

		}
		pre := i + 1
		tail := len(nums) - 1
		for pre < tail {
			res := nums[i] + nums[pre] + nums[tail] - target
			if res == 0 {
				return target
			} else if res > 0 {
				if abs(sum-target) > res {
					sum = nums[i] + nums[pre] + nums[tail]
				}
				tail--
			} else {
				if abs(sum-target) > abs(res) {
					sum = nums[i] + nums[pre] + nums[tail]
				}
				pre++
			}
		}
	}
	return sum
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
