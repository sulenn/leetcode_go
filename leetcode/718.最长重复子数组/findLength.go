package main

//https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/

//思路：暴力，依次遍历所有
//func findLength(A []int, B []int) int {
//	max := 0
//	for i:=0; i<len(A); i++ {
//		for j:=0; j<len(B); j++ {
//			if A[i] == B[j] {
//				num_i := i+1
//				num_j := j+1
//				for num_i<len(A) && num_j<len(B) && A[num_i] == B[num_j] {
//					num_i++
//					num_j++
//				}
//				if num_i - i > max {  // 最大值比较
//					max = num_i - i
//				}
//			}
//			if len(B)-j-1 <= max {   // 剪枝
//				break
//			}
//		}
//		if len(A)-i-1 <= max {   // 剪枝
//			break
//		}
//	}
//	return max
//}

// 动态规划
func findLength(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		dp[i] = make([]int, len(B)+1)
	}
	max := 0
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}
