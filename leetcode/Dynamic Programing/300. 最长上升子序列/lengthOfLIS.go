package main

//https://leetcode-cn.com/problems/longest-increasing-subsequence/

//参考：https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/guan-fang-ti-jie-kan-bu-dong-de-kan-zhe-li-golang-/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, res := make([]int, len(nums)), 0
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
