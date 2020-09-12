package main

//https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	height := len(grid)
	width := len(grid[0])
	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			if i < height-1 && j == width-1 {
				grid[i][j] += grid[i+1][j]
			}
			if j < width-1 && i == height-1 {
				grid[i][j] += grid[i][j+1]
			}
			if i < height-1 && j < width-1 {
				if grid[i+1][j] > grid[i][j+1] {
					grid[i][j] += grid[i+1][j]
				} else {
					grid[i][j] += grid[i][j+1]
				}
			}
		}
	}
	return grid[0][0]
}
