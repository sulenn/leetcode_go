package main

//https://leetcode-cn.com/problems/as-far-from-land-as-possible/

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	land := [][2]int{}   // 注意这里[2]，限制数组中每一个元素长度为2
	visited := make([][]int, m)   // 用于记录节点是否已被访问
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				visited[i][j] = 1
				land = append(land, [2]int{i, j})
			}
		}
	}
	if len(land) == 0 || len(land) == m * n {  // 全0 or 全1
		return -1
	}
	start, end := 0, len(land)
	r := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}  // 上下左右
	for start < end {   // 精彩的写法
		x, y := land[start][0], land[start][1]
		start++
		for _, r1 := range r {
			x1, y1 := x + r1[0], y + r1[1]
			if x1 < 0 || x1 >= len(grid) || y1 < 0 || y1 >= len(grid[0]) {
				continue
			}
			if visited[x1][y1] > 0 {
				continue
			}
			visited[x1][y1] = visited[x][y] + 1   // 记录当前节点的访问层级。
			land = append(land, [2]int{x1, y1})
			end++
		}
	}
	max := -1
	for i := 0; i < m; i++ {  // 精彩的写法
		for j := 0; j < n; j++ {
			if visited[i][j] > max {
				max = visited[i][j]
			}
		}
	}
	return max - 1
}