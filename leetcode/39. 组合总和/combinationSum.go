package main

import "fmt"

//https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtracking(candidates, target, &result, []int{})
	return result
}

func backtracking(candidates []int, target int, combinations *[][]int, curNums []int) {
	if target == 0 {
		deepcopy := make([]int, len(curNums))
		copy(deepcopy, curNums)
		*combinations = append(*combinations, deepcopy)
		return
	}
	if target < 0 {
		return
	}
	for k, v := range candidates {
		backtracking(candidates[k:], target-v, combinations, append(curNums, v))
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
