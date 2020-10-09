package main

//https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	numDic := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := numDic[nums[i]]; ok {
			return []int{v, i}
		}
		numDic[target-nums[i]] = i
	}
	return []int{}
}
