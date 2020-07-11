package main

import "fmt"

//https://leetcode-cn.com/problems/number-of-islands/

//广搜
func numIslands(grid [][]byte) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				islandList := [][]int{{i, j}}
				grid[i][j] = '0'
				for len(islandList) > 0 {
					island := islandList[0]
					islandList = islandList[1:]
					if island[1]+1 < len(grid[0]) && grid[island[0]][island[1]+1] == '1' { //右
						islandList = append(islandList, []int{island[0], island[1] + 1})
						grid[island[0]][island[1]+1] = '0' // 如果不设置 ’0‘，会导致超时，下同
					}
					if island[1]-1 >= 0 && grid[island[0]][island[1]-1] == '1' { //左
						islandList = append(islandList, []int{island[0], island[1] - 1})
						grid[island[0]][island[1]-1] = '0'
					}
					if island[0]-1 >= 0 && grid[island[0]-1][island[1]] == '1' { // 上
						islandList = append(islandList, []int{island[0] - 1, island[1]})
						grid[island[0]-1][island[1]] = '0'
					}
					if island[0]+1 < len(grid) && grid[island[0]+1][island[1]] == '1' { // 下
						islandList = append(islandList, []int{island[0] + 1, island[1]})
						grid[island[0]+1][island[1]] = '0'
					}
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
}
