package main

import "sort"

//https://leetcode-cn.com/problems/combination-sum-ii/description/

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	result := &[][]int{}
	backtracking(candidates, target, result, []int{}, 0)
	return *result
}

func backtracking(candidates []int, target int, combinations *[][]int, curNums []int, curSum int) {
	if target == curSum { // 深拷贝
		deepcopy := make([]int, len(curNums))
		copy(deepcopy, curNums)
		*combinations = append(*combinations, deepcopy)
		return
	}
	if curSum > target {
		return
	}
	for k, v := range candidates {
		if k != 0 && candidates[k] == candidates[k-1] {
			continue
		}
		backtracking(candidates[k+1:], target, combinations, append(curNums, v), curSum+v) // 注意前K-1个元素不能出现在下一次递归中
	}
}
