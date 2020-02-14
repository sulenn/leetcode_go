package main

import "errors"

//https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/

//二分法
func findMin(nums []int) int {
	if len(nums) == 0 {
		panic(errors.New("数组为空！"))
	}
	if nums[0] <= nums[len(nums)-1] {   // 针对 [1、2] 这种特殊的测试用例
		return nums[0]
	}
	prev := 0
	tail := len(nums) - 1
	for tail >= prev {
		mid := (prev + tail) / 2
		if mid > 0 && nums[mid-1] > nums[mid] {   // nums[mid] < nums[mid-1]
			return nums[mid]
		} else if mid < len(nums) && nums[mid+1] < nums[mid] {   // nums[mid+1] < nums[mid]
			return nums[mid+1]
		}
		if nums[mid] > nums[0] {
			prev = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return 0
}

