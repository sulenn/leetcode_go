package main

import "sort"

//https://leetcode-cn.com/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	backTracking(nums, []int{}, &result)
	return result
}

func backTracking(nums []int, curNums []int, result *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(curNums))
		copy(temp, curNums)
		*result = append(*result, temp)
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		newNums := make([]int, len(nums)-1)
		copy(newNums[:i], nums[:i])
		copy(newNums[i:], nums[i+1:])
		backTracking(newNums, append(curNums, nums[i]), result)
	}
}
