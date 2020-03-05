package main

//https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/

// 双指针
func twoSum(nums []int, target int) []int {
	start := 0
	end := len(nums)-1
	for start < end {
		if nums[start] + nums[end] == target {
			return []int {nums[start], nums[end]}
		}
		if nums[start] + nums[end] > target {
			end--
		} else {
				start++
		}
	}
	return []int {}
}