package main

//https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

// 双指针
func findContinuousSequence(target int) [][]int {
	result := [][]int{}
	start := 1
	end := 2
	for start < end {
		if (start+end)*(end-start+1)/2 == target { // 等差求和
			subResult := []int{}
			for i := start; i <= end; i++ { // 满足条件的连续和
				subResult = append(subResult, i)
			}
			result = append(result, subResult)
			start++
		}
		if (start+end)*(end-start+1)/2 > target {
			start++
		} else {
			end++
		}
	}
	return result
}
