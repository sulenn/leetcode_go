package main

//https://leetcode-cn.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	newNums := make([]int, 0)
	for i := 0; i < n; i++ {
		first := i
		second := n + i
		newNums = append(newNums, nums[first], nums[second])
	}
	return newNums
}
