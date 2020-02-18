package main

import "fmt"

//https://leetcode-cn.com/problems/unique-paths-ii/

//该题和 Unique Path 以及 Minimun Path Sum 的解法一样，动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {   // 异常情况
		return 0
	}
	m := len(obstacleGrid)   // 行
	n := len(obstacleGrid[0])   // 列
	for i := 0; i< m; i++{
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else if i-1 >= 0 && j-1 >= 0 {   // 上、左都在
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else if i-1 >= 0 && j-1 < 0 {   // 上在
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
			} else if i-1 < 0 && j-1 >=0 {   // 左在
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			} else {    // 起点
				obstacleGrid[i][j] = 1
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int {{0,0,0},{0,1,0},{0,0,0}}))
	fmt.Println(uniquePathsWithObstacles([][]int {{0}}))
}
