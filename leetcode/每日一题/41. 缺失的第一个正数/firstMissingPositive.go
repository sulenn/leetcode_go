package main

//https://leetcode-cn.com/problems/first-missing-positive/

// 思路：原地按下标排序，利用数据本身空间
func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; { // 原地桶排序
		if nums[i]-1 > i {
			if nums[i]-1 < length && nums[nums[i]-1] != nums[i] {
				nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
				continue
			}
		} else if nums[i] > 0 {
			nums[nums[i]-1] = nums[i]
		}
		i++
	}
	for i := 0; i < length; i++ { // 查找缺失的最小正整数
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}
