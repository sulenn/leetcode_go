package main

//dp[i] = max(dp[i-1], dp[i-2]+sort[i]*i)
func deleteAndEarn(nums []int) int {
	sort := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		sort[nums[i]]++
	}
	dp := make([]int, 10001)
	dp[1] = sort[1] * 1
	for i := 2; i <= 10000; i++ {
		if dp[i-1] > dp[i-2]+sort[i]*i {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + sort[i]*i
		}
	}
	return dp[10000]
}
