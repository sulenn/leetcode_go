package main

import "fmt"

//https://leetcode-cn.com/problems/minimum-path-sum/

//dp
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rowLen := len(grid) - 1
	colLen := len(grid[0]) - 1
	for i := rowLen; i >= 0; i-- { // 倒序求每个点的路径值，从最下一行开始
		for j := colLen; j >= 0; j-- {
			if j+1 <= colLen && i+1 <= rowLen {
				if grid[i][j+1] > grid[i+1][j] {
					grid[i][j] += grid[i+1][j]
				} else {
					grid[i][j] += grid[i][j+1]
				}
			} else if j+1 <= colLen && i+1 > rowLen {
				grid[i][j] += grid[i][j+1]
			} else if j+1 > colLen && i+1 <= rowLen {
				grid[i][j] += grid[i+1][j]
			}
		}
	}
	return grid[0][0]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
