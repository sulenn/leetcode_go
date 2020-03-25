package main

//https://leetcode-cn.com/problems/surface-area-of-3d-shapes/

func surfaceArea(grid [][]int) int {
	if len(grid) == 0 {return 0}
	result := 0
	for i:=0 ;i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {  // 每一个坐标靠考虑左边和右边以及上下重叠
			if i == 0 && j != 0 {  // 第一行
				if grid[i][j] < grid[i][j-1] {  //左边
					result -= grid[i][j]*2
				} else {
					result -= grid[i][j-1]*2
				}
			} else if j == 0 && i != 0 {  // 第一列
				if grid[i][j] < grid[i-1][j] {  //下边
					result -= grid[i][j]*2
				} else {
					result -= grid[i-1][j]*2
				}
			} else if i != 0 && j != 0 {  // 非第一行、第一列
				if grid[i][j] < grid[i][j-1] { // 左边
					result -= grid[i][j]*2
				} else {
					result -= grid[i][j-1]*2
				}
				if grid[i][j] < grid[i-1][j] { // 右边
					result -= grid[i][j]*2
				} else {
					result -= grid[i-1][j]*2
				}
			}
			if grid[i][j] != 0 {
				result += grid[i][j]*6-(grid[i][j]-1)*2
			}

		}
	}
	return result
}