package main

import "sort"

//https://leetcode-cn.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {return [][]int {}}
	sort.Slice(intervals, func(i, j int) bool {   // 按区间头部排序
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	interval := intervals[0]  // 当前累计区间
	for i:=1; i<len(intervals); i++ {
		if intervals[i][0] >= interval[0] && intervals[i][0] <= interval[1] {
			if intervals[i][1] > interval[1] {
				interval[1] = intervals[i][1]
			}
		} else {
			result = append(result, interval)
			interval = intervals[i]
		}
	}
	result = append(result, interval)
	return result
}