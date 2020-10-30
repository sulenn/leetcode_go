package main

//https://leetcode-cn.com/problems/island-perimeter/

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	column := len(grid[0])
	result := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				result += 4
				if i+1 < row && grid[i+1][j] == 1 {
					result -= 2
				}
				if j+1 < column && grid[i][j+1] == 1 {
					result -= 2
				}
			}
		}
	}
	return result
}
