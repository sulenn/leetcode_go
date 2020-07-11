package main

import "sort"

//https://leetcode-cn.com/problems/non-overlapping-intervals/description/

func eraseOverlapIntervals(intervals [][]int) int {
	//for i:=0; i<len(intervals); i++ {   // 插入排序，先比较区间末端，如果末端相等，再比较区间头
	//	for j:=i;j>0;j-- {
	//		if intervals[j][1] < intervals[j-1][1] || (intervals[j][1] == intervals[j-1][1] && intervals[j-1][0] < intervals[j][0]) {
	//			intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
	//		}
	//	}
	//}
	sort.Slice(intervals, func(i, j int) bool { // 排序，只考虑区间末端即可
		return intervals[i][1] < intervals[j][1]
	})
	count := 0
	cur := []int{} // 最近的区间
	for j := 0; j < len(intervals); j++ {
		if j == 0 {
			count++
			cur = intervals[0]
			continue
		}
		if intervals[j][0] >= cur[1] {
			cur = intervals[j]
			count++
		}
	}
	return len(intervals) - count
}
