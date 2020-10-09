package main

//https://leetcode-cn.com/problems/predict-the-winner/

// 参考：https://leetcode-cn.com/problems/predict-the-winner/solution/shou-hua-tu-jie-san-chong-xie-fa-di-gui-ji-yi-hua-/
func PredictTheWinner(nums []int) bool {
	length := len(nums)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = nums[i]
	}
	var sum int
	for i := length - 2; i >= 0; i-- {
		sum = nums[i]
		for j := i + 1; j < length; j++ {
			sum += nums[j]
			left := sum - dp[i][j-1]
			right := sum - dp[i+1][j]
			if left > right {
				dp[i][j] = left
			} else {
				dp[i][j] = right
			}
		}
	}
	return 2*dp[0][length-1] >= sum
}
