package main

import "math"

//https://leetcode-cn.com/problems/triangle/

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	for i := 1; i < rows; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
				continue
			}
			if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
				continue
			}
			if triangle[i-1][j] < triangle[i-1][j-1] {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += triangle[i-1][j-1]
			}
		}
	}
	min := math.MaxInt64
	for i := 0; i < len(triangle[rows-1]); i++ {
		if triangle[rows-1][i] < min {
			min = triangle[rows-1][i]
		}
	}
	return min
}
