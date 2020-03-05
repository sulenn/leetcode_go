package main

import "sort"

//https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/

func isStraight(nums []int) bool {
	sort.Ints(nums)
	sum0 := 0 // 数组中0的数量
	for i:=0; i<len(nums); i++ {
		if nums[i] == 0 {
			sum0++
		} else {break}
	}
	sum1 := 0 // 不连续的数量
	for i:=sum0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {return false}  //存在对子
		sum1 += nums[i+1] - nums[i] - 1
	}
	if sum0 >= sum1 {
		return true
	}
	return false
}

