package main

import "math"

//https://leetcode-cn.com/problems/arithmetic-slices/

func numberOfArithmeticSlices(A []int) int {
	result := 0
	num := 0  // 等差数列中数字数量
	value := math.MaxInt64  // 差值
	for i:=0; i< len(A)-1; i++ {
		if A[i+1] - A[i] == value {
			num++
			continue
		} else if num >= 3 {
			result += (1+num-2)*(num-2)/2    // 利用等差数列求和计算满足要求的子数组；也可以用递归的方法来计算
		}
		num = 2
		value = A[i+1] - A[i]
	}
	if num >= 3 {
		result += (1+num-2)*(num-2)/2    // 利用等差数列求和计算满足要求的子数组；也可以用递归的方法来计算
	}
	return result
}
