package main

//https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	freq := make([]int, 101)
	result := make([]int, len(nums))
	// 统计出现频率 frequency
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}
	// 对频率(而非对原数组nums)从前到后累加
	for i := 1; i < 101; i++ {
		freq[i] += freq[i-1]
	}
	// 输出结果
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result[i] = freq[nums[i]-1]
		}
	}
	return result
}
