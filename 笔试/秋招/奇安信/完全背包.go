package main

import (
	"fmt"
	"io"
)

func main() {
	var target int
	var count int
	for {
		_, err := fmt.Scan(&target, &count)
		if err == io.EOF {
			break
		}
		nums := make([][2]int, count)
		for i := 0; i < count; i++ {
			fmt.Scan(&nums[i][0], &nums[i][1])
		}
		dp := make([]int, target+1)
		for i := 0; i < len(nums); i++ {
			for j := nums[i][0]; j <= target; j++ {
				dp[j] = max(dp[j], dp[j-nums[i][0]]+nums[i][1])
			}
		}
		fmt.Println(dp[target])
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
