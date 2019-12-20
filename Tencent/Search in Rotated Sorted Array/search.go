package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

// 留意存在未旋转的数组，如[1]、[1,3] 等
func search(nums []int, target int) int {
	if len(nums) <= 0 { // 排除异常
		return -1
	}
	pivot := findPivot(nums)
	left := findTarget(nums[:pivot], target)
	if left != -1 {
		return left
	}
	right := findTarget(nums[pivot:], target)
	if right != -1 {
		return right + pivot
	}
	return -1
}
//排序数组中查找目标值
func findTarget(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {   // 存在等于号
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
//查找旋转点
func findPivot(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {  //留意没有等于号
		mid := (left + right) / 2
		if mid > 0 && nums[mid] < nums[mid - 1] {
			return mid
		}
		if nums[mid] > nums[left] {
			left = mid
		} else {
			right = mid
		}
	}
	return len(nums) - 1   // 注意，存在没有旋转的排序数组
}

func main() {
	fmt.Println(search([]int {4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int {4,5,6,7,0,1,2}, 3))
	fmt.Println(search([]int {1}, 1))
	fmt.Println(search([]int {1, 3}, 1))
}
