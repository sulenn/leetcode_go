package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var n1, n2 int
	for {
		_, err := fmt.Scan(&n1, &n2)
		if err == io.EOF {
			break
		}
		x := make([]int, n1+n2)
		dp := make([]int, n1+n2)
		for i := 0; i < n1+n2; i++ {
			fmt.Scan(&x[i])
			dp[i] = 1
		}
		for i := 1; i < n1+n2; i++ {
			for j := 0; j < i; j++ {
				if x[i] >= x[j] && dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		max := math.MinInt64
		for i := 0; i < n1+n2; i++ {
			if dp[i] > max {
				max = dp[i]
			}
		}
		fmt.Println(max)
	}
}
