package main

import (
	"fmt"
	"io"
)

func main() {
	var n, m int
	for {
		_, err := fmt.Scan(&n, &m)
		if err == io.EOF {
			break
		}
		nums := make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&nums[i][0], &nums[i][1])
		}
		dp := make([]int, m+1)
		for i := nums[0][0]; i <= m; i++ {
			dp[i] = nums[0][1]
		}
		for i := 1; i < n; i++ {
			for j := m; j >= nums[i][0]; j-- {
				if nums[i][1]+dp[j-nums[i][0]] > dp[j] {
					dp[j] = nums[i][1] + dp[j-nums[i][0]]
				}
			}
		}
		fmt.Println(dp[m])
	}
}
