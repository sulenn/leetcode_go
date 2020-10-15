package main

//https://leetcode-cn.com/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) int {
	result := 0
	sums := [101]int{}
	for i := 0; i < len(nums); i++ {
		result += sums[nums[i]]
		sums[nums[i]]++
	}
	return result
}
