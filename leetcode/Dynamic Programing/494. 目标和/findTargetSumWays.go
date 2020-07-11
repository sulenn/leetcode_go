package main

import "math"

//https://leetcode-cn.com/problems/target-sum/description/

//0/1 背包问题，参考：https://leetcode-cn.com/problems/target-sum/solution/dong-tai-gui-hua-si-kao-quan-guo-cheng-by-keepal/
func findTargetSumWays(nums []int, S int) int {
	arr := make([][]int, len(nums)+1)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if math.Abs(float64(sum)) < math.Abs(float64(S)) {
		return 0
	}
	for i := 0; i < len(nums)+1; i++ {
		arr[i] = make([]int, 2*sum+1)
	}
	arr[0][sum] = 1
	for i := 1; i < len(nums)+1; i++ {
		for j := 0; j < 2*sum+1; j++ {
			if j+nums[i-1] < 2*sum+1 {
				arr[i][j] += arr[i-1][j+nums[i-1]]
			}
			if j-nums[i-1] >= 0 {
				arr[i][j] += arr[i-1][j-nums[i-1]]
			}
		}
	}
	return arr[len(nums)][sum+S]
}
