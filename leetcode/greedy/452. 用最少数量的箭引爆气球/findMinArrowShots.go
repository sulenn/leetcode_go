package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/description/

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {return 0}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	end := points[0][1]
	for i:=1; i<len(points); i++ {
		if points[i][0] <= end {
			continue
		}
		count++
		end = points[i][1]
	}
	return count
}
