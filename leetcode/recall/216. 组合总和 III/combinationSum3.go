package main

//https://leetcode-cn.com/problems/combination-sum-iii/

func combinationSum3(k int, n int) [][]int {
	if n > 55 {
		return [][]int{}
	}
	result := &[][]int{}
	backtracking(1, k, n, result, []int{})
	return *result
}

func backtracking(start int, k int, n int, combinations *[][]int, curNums []int) {
	if n == 0 && len(curNums) == k { // 深拷贝
		deepcopy := make([]int, len(curNums))
		copy(deepcopy, curNums)
		*combinations = append(*combinations, deepcopy)
		return
	}
	if n < 0 {
		return
	}
	for i := start; i <= 9; i++ {
		backtracking(i+1, k, n-i, combinations, append(curNums, i)) // 注意前K-1个元素不能出现在下一次递归中
	}
}
