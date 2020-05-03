package main

import "math"

//https://leetcode-cn.com/problems/maximum-product-subarray/

//maxDP[i + 1] = max(maxDP[i] * A[i + 1], A[i + 1],minDP[i] * A[i + 1])
//minDP[i + 1] = min(minDP[i] * A[i + 1], A[i + 1],maxDP[i] * A[i + 1])
//dp[i + 1] = max(dp[i], maxDP[i + 1])
//参考：https://leetcode-cn.com/problems/maximum-product-subarray/solution/dpfang-fa-xiang-jie-by-yang-cong-12/
func maxProduct(nums []int) int {
	max := math.MinInt64
	imax := 1   // 当前最大值
	imin := 1   // 当前最小值
	for i:=0; i< len(nums); i++ {
		imax, imin = MaxAndMin(imax*nums[i], nums[i], imin*nums[i])
		if max < imax {max = imax}
	}
	return max
}

func MaxAndMin(a,b,c int) (int,int) {
	if a >= b && b >= c {
		return a,c
	}
	if a >= c && c >= b {
		return a,b
	}
	if b >= a && a >= c {
		return b, c
	}
	if b >= c && c >= a {
		return b, a
	}
	if c >= a && a >= b {
		return c, b
	}
	if c >= b && c >= a {
		return c, a
	}
	return 0 , 0
}

