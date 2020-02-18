package main

import "fmt"

//https://leetcode-cn.com/problems/island-perimeter/

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	sum1 := 0    // 1的方格数量
	sum2 := 0     // 相邻的方格数量（两个方格相邻算1）
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sum1++
				if j+1 <len(grid[0]) && grid[i][j+1] == 1{   // 右相邻
					sum2++
				}
				if i+1 < len(grid) && grid[i+1][j] == 1{    // 下相邻
					sum2++
				}
			}
		}
	}
	return sum1*4 - sum2*2    // 周长公式
}

func main() {
	fmt.Println(islandPerimeter([][]int {{0,1,0,0}, {1,1,1,0},{0,1,0,0},{1,1,0,0}}))
}
