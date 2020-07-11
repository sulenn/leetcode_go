package main

import "sort"

//https://leetcode-cn.com/problems/video-stitching/

//贪心，参考：https://leetcode-cn.com/problems/video-stitching/solution/tan-xin-suan-fa-c-by-f821878167/
func videoStitching(clips [][]int, T int) int {
	sort.Slice(clips, func(i, j int) bool { // 升序排列
		return clips[i][0] < clips[j][0]
	})
	max := 0   // 每回合区间的尾端值
	count := 0 // 代取的区间数目
	i := 0     // 数组下标
	for max < T && i < len(clips) {
		t := max // 满足条件的当前回合区间尾端值
		for j := i; j < len(clips) && max >= clips[j][0]; j++ {
			if t < clips[j][1] {
				t = clips[j][1]
			}
			i = j // 数组下标累加
		}
		if t != max {
			count++
			max = t
		} else {
			break
		}
	}
	if max >= T {
		return count
	}
	return -1
}
