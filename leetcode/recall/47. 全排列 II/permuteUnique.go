package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/permutations-ii/

//回溯，方法可参考 - https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E6%90%9C%E7%B4%A2.md#1-%E6%95%B0%E5%AD%97%E9%94%AE%E7%9B%98%E7%BB%84%E5%90%88
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums) // 排序
	result := &[][]int{}
	backtracking(nums, result, []int{})
	return *result
}

func backtracking(nums []int, combinations *[][]int, curNums []int) {
	if len(nums) == 0 {
		*combinations = append(*combinations, curNums)
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		temp_curNums := make([]int, len(curNums)) // 硬拷贝，不然会出错
		copy(temp_curNums, curNums)
		temp_curNums = append(temp_curNums, nums[i])
		temp_nums := make([]int, len(nums)) // 硬拷贝。这里也可以用一个visited数组来做标记。
		copy(temp_nums, nums)
		temp_nums = append(temp_nums[:i], temp_nums[i+1:]...)
		backtracking(temp_nums, combinations, temp_curNums)
	}
}
