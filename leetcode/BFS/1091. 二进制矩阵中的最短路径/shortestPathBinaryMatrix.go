package main

//https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 {return -1}
	length := len(grid)
	cur := [][]int {{0,0}}
	location := [][]int {{0,-1},{1,-1},{1,0},{1,1},{0,1},{-1,1},{-1,0},{-1,-1}}   // 8个方向
	depth := 0
	for len(cur) != 0 {
		next := make([][]int, 0)
		depth++
		for i:=0; i < len(cur); i++ {
			if grid[cur[i][0]][cur[i][1]] == 1 {
				continue
			}
			if cur[i][0] == length-1 && cur[i][1] == length-1 {return depth}
			grid[cur[i][0]][cur[i][1]] = 1
			for j:=0; j<8;j++ {
				x := cur[i][0] + location[j][0]
				y := cur[i][1] + location[j][1]
				if x >=0 && x <length && y >=0 && y <length && grid[x][y] != 1 {
					next = append(next, []int{x,y})
				}
			}
		}
		cur = next
	}
	return -1
}
