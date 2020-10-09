package main

// 此外可参考单调栈解法O(n)：https://leetcode-cn.com/problems/daily-temperatures/solution/leetcode-tu-jie-739mei-ri-wen-du-by-misterbooo/

//带剪枝的暴力，时间复杂度O(n^2)
func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	max := 0
	for i := len(T) - 1; i >= 0; i-- {
		if T[i] >= max {
			result[i] = 0
			max = T[i]
		} else {
			for j := i + 1; j < len(T); j++ {
				if T[j] > T[i] {
					result[i] = j - i
					break
				}
			}
		}
	}
	return result
}
