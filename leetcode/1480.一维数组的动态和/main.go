package main

//https://leetcode-cn.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	preSum := 0
	for i := 0; i < len(nums); i++ {
		nums[i] += preSum
		preSum = nums[i]
	}
	return nums
}
