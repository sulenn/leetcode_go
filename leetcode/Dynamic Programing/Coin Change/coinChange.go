package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/coin-change/

//DP,可以参考解法：https://leetcode-cn.com/problems/coin-change/solution/dong-tai-gui-hua-tao-lu-xiang-jie-by-wei-lai-bu-ke/
func coinChange(coins []int, amount int) int {
	if amount == 0 { // 总额为 0，直接输出 0
		return 0
	}
	size := len(coins) // 没有硬币
	if size == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	/*
	   dp[i] 达到 i 的钱数所需要的最少的硬币数
	   dp[i] = min(dp[i], dp[i-coin]+1)
	*/
	for i := 0; i <= amount; i++ {
		for j := 0; j < size; j++ {
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//func coinChange(coins []int, amount int) int {
//	if amount == 0 {   // 总额为 0，直接输出 0
//		return 0
//	}
//	minCoinNumsList := make(map[int]int)
//	for _,coin := range coins {    // 初始化
//		if amount == coin {   // 恰巧 ammout 等于某一种硬币的币值
//			return 1
//		} else {
//			minCoinNumsList[coin] = 1
//		}
//	}
//	i := 1
//	//硬币可能是乱序的， 不能用下面这个方法
//	//for i <= amount {
//	//	for j:=len(coins)-1; j>=0; j-- {
//	//		if value, ok := minCoinNumsList[i-coins[j]]; ok {
//	//			if _, ok := minCoinNumsList[i]; !ok {
//	//				minCoinNumsList[i] = value + 1
//	//			}
//	//		}
//	//	}
//	//	i++
//	//}
//	for i <= amount {
//		for _, coin := range coins {
//			if value1, ok := minCoinNumsList[i-coin]; ok {
//				if value2 ,ok := minCoinNumsList[i]; ok {
//					if value1 + 1 < value2 {
//						minCoinNumsList[i] = value1 + 1
//					}
//				} else {
//					minCoinNumsList[i] = value1 + 1
//				}
//			}
//		}
//		i++
//	}
//	if value, ok := minCoinNumsList[amount]; ok {
//		return value
//	} else {
//		 return -1
//	}
//}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 1))
}
