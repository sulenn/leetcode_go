package main

//https://leetcode-cn.com/problems/subsets/

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	backtracking(nums, &result, []int{})
	return result
}

func backtracking(nums []int, result *[][]int, curArr []int) {
	tempCurArr := make([]int, len(curArr))
	copy(tempCurArr, curArr)
	*result = append(*result, tempCurArr)
	for i := 0; i < len(nums); i++ {
		backtracking(nums[i+1:], result, append(curArr, nums[i]))
	}
}
