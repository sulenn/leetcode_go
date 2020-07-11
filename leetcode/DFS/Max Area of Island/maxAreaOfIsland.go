package main

import "fmt"

//https://leetcode-cn.com/problems/max-area-of-island/

// 广搜
func maxAreaOfIsland(grid [][]int) int {
	result := 0
	var dfs func(grid [][]int, i int, j int)
	sum := 0
	dfs = func(grid [][]int, i int, j int) {
		grid[i][j] = 0
		sum++
		if j-1 >= 0 && grid[i][j-1] == 1 {
			dfs(grid, i, j-1)
		} // 上
		if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
			dfs(grid, i, j+1)
		} //下
		if i-1 >= 0 && grid[i-1][j] == 1 {
			dfs(grid, i-1, j)
		} //左
		if i+1 < len(grid) && grid[i+1][j] == 1 {
			dfs(grid, i+1, j)
		} //右
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sum = 0
				dfs(grid, i, j)
				if sum > result { //  交换最大岛屿面积
					result = sum
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}))
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}))
}
