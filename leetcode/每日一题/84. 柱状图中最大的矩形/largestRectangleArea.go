package main

//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

// 单调栈，参考：https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/84-by-ikaruga/
//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/zhao-liang-bian-di-yi-ge-xiao-yu-ta-de-zhi-by-powc/
func largestRectangleArea(heights []int) int {
	max := 0
	stack := make([]int, 0)
	heights = append(heights, 0) // 注意填充0，避免数组单调递增，没有出口
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			cur := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // 删掉栈中最后一位数
			left := 0
			if len(stack) > 0 { // 小心栈溢出
				left = stack[len(stack)-1] + 1
			}
			right := i - 1
			if max < (right-left+1)*cur {
				max = (right - left + 1) * cur
			}
		}
		stack = append(stack, i)
	}
	return max
}

// O(n^2) 时间复杂度，暴力
//func largestRectangleArea(heights []int) int {
//	max := 0
//	for i:=0; i<len(heights); i++ {
//		if heights[i] == 0 {continue}
//		minValue := heights[i]    // 当前遍历批次中最短的列
//		for j:=i; j<len(heights); j++ {
//			if heights[j] == 0 {break}
//			if heights[j] < minValue {minValue = heights[j]}
//			if (j-i+1)*minValue > max {max = (j-i+1)*minValue }
//		}
//	}
//	return max
//}
