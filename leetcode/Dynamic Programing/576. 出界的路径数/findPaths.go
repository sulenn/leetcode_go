package main

//https://leetcode-cn.com/problems/out-of-boundary-paths/

func findPaths(m int, n int, N int, i int, j int) int {
	ballMoveArr := make([][]int, m)
	for x:=0; x<m; x++ {
		ballMoveArr[x] = make([]int, n)
	}
	ballMoveArr[i][j] = 1   // 初始点
	count := 0
	for x:=0; x<N; x++ {  // N 次以内可以出边界的路径总数
		count += calCount(ballMoveArr)
		count %= 1000000007    // 注意取余
		ballMoveArr = move(ballMoveArr)
	}
	return count
}

func move(ballMoveArr [][]int) [][]int {   // 移动数组中的球
	r := len(ballMoveArr)
	c := len(ballMoveArr[0])
	temp := make([][]int, r)   // balMoveArr 进行移动
	for x:=0; x<r; x++ {
		temp[x] = make([]int, c)
	}
	location := [][]int {{0, -1}, {0, 1}, {1, 0}, {-1, 0}}   // 上下左右
	for i:=0; i<r; i++ {
		for j:=0; j<c; j++ {
			for v:=0; v<4; v++ {
				x := i+location[v][0]
				y := j+location[v][1]
				if x >= 0 && x < r && y >= 0 && y < c {  // 注意取余
					temp[x][y] += ballMoveArr[i][j]
					temp[x][y] %= 1000000007
				}
			}
		}
	}
	return temp
}

func calCount(ballMoveArr [][]int) int {   // 计算当前数组中可移动出边界的次数
	r := len(ballMoveArr)
	c := len(ballMoveArr[0])
	count := 0
	for i:=0; i<c; i++ {  // 第一行和最后一行
		count += ballMoveArr[0][i]
		count += ballMoveArr[r-1][i]
	}

	for i:=0; i<r; i++ {  // 第一列和最后一列
		count += ballMoveArr[i][0]
		count += ballMoveArr[i][c-1]
	}
	return count
}
