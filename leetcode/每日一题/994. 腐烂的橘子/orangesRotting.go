package main

//https://leetcode-cn.com/problems/rotting-oranges/

func orangesRotting(grid [][]int) int {
	result := 0
	badOrangePerTime := [][]int{}
	for i := 0; i < len(grid); i++ { // 第一次遍历数组中所有腐烂的橘子
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				badOrangePerTime = append(badOrangePerTime, []int{i, j})
			}
		}
	}
	for len(badOrangePerTime) > 0 { // 腐烂的句子感染腐烂的橘子
		badOrangeNextTime := [][]int{}
		for i := 0; i < len(badOrangePerTime); i++ {
			curLacation := badOrangePerTime[i]
			if curLacation[0]-1 >= 0 && grid[curLacation[0]-1][curLacation[1]] == 1 { // 上
				grid[curLacation[0]-1][curLacation[1]] = 2
				badOrangeNextTime = append(badOrangeNextTime, []int{curLacation[0] - 1, curLacation[1]})
			}
			if curLacation[0]+1 < len(grid) && grid[curLacation[0]+1][curLacation[1]] == 1 { // 下
				grid[curLacation[0]+1][curLacation[1]] = 2
				badOrangeNextTime = append(badOrangeNextTime, []int{curLacation[0] + 1, curLacation[1]})
			}
			if curLacation[1]-1 >= 0 && grid[curLacation[0]][curLacation[1]-1] == 1 { // 左
				grid[curLacation[0]][curLacation[1]-1] = 2
				badOrangeNextTime = append(badOrangeNextTime, []int{curLacation[0], curLacation[1] - 1})
			}
			if curLacation[1]+1 < len(grid[0]) && grid[curLacation[0]][curLacation[1]+1] == 1 { // 右
				grid[curLacation[0]][curLacation[1]+1] = 2
				badOrangeNextTime = append(badOrangeNextTime, []int{curLacation[0], curLacation[1] + 1})
			}
		}
		if len(badOrangeNextTime) != 0 { // 如果本次确定右感染的橘子，次数+1
			result++
		}
		badOrangePerTime = badOrangeNextTime
	}
	for i := 0; i < len(grid); i++ { // 橘子感染完之后查看有没有独立的橘子未被感染
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return result
}
