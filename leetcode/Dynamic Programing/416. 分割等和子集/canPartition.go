package main

//https://leetcode-cn.com/problems/partition-equal-subset-sum/

//0/1 背包问题，dp[i][j] = dp[i][j-1] || dp[i][j-v0]
func canPartition(nums []int) bool {
	sums := 0
	for _, v:=range nums {
		sums += v
	}
	if sums == 0 ||sums % 2 == 1 {return false}
	arr := make([]bool, sums/2+1)
	arr[0]= true
	for i:=0; i<len(nums); i++ {
		for j:=sums/2; j>i; j-- {
			if j-nums[i] >= 0 {
				arr[j] = arr[j] || arr[j-nums[i]]
			} else {
				arr[j] = arr[j]
			}
		}
	}
	//二维数组方法
	//for i:=0; i < len(arr); i++ {
	//	arr[i] = make([]bool, sums/2+1)
	//}
	//arr[0][0] = true
	//for i:=1; i<len(arr); i++ {
	//	for j:=1; j<sums/2+1; j++ {
	//		if j-nums[i-1] >= 0 {
	//			arr[i][j] = arr[i-1][j] || arr[i-1][j-nums[i-1]]
	//		} else {
	//			arr[i][j] = arr[i-1][j]
	//		}
	//	}
	//}
	return arr[sums/2]
}