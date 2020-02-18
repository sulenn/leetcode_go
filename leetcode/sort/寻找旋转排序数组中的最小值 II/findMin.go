package main

//https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

import "errors"

//二分法
func findMin(nums []int) int {
	if len(nums) == 0 {
		panic(errors.New("数组为空！"))
	}
	if nums[0] < nums[len(nums)-1] {   // 针对 [1、2] 这种特殊的测试用例
		return nums[0]
	}
	prev := 0
	tail := len(nums) - 1
	for tail > prev {
		mid := (prev + tail) / 2
		if mid-1 >= 0 && nums[mid-1] > nums[mid] {   // nums[mid] < nums[mid-1]
			return nums[mid]
		} else if mid+1 < len(nums) && nums[mid+1] < nums[mid] {   // nums[mid+1] < nums[mid]
			return nums[mid+1]
		}
		if nums[mid] == nums[prev] && nums[mid] == nums[tail] {   // 当 prev、mid、tail 的值相等时无法判断旋转点在何处。所以用顺序查找
			return inOrder(nums, prev, tail)
		}
		if nums[mid] >= nums[prev] {
			prev = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return nums[0]
}

func inOrder(nums []int, prev int, tail int) int {  // 按顺序搜索
	for ;prev < tail; prev++ {
		if nums[prev] > nums[prev+1] {
			return nums[prev+1]
		}
	}
	return nums[prev]
}