package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	curSum := nums[0] // 当前最大连续子序列和
	maxSum := nums[0] // 全局最大连续子序列和
	for _, elem := range nums[1:] {
		if curSum+elem >= elem {
			curSum += elem
		} else {
			curSum = elem
		}
		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
