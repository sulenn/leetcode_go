package main

//https://leetcode-cn.com/problems/surrounded-regions/description/

//思路：先处理矩阵外围四行，存在'O'，则深搜将所有临界的'O'都变为'W'
// 之后在遍历一遍矩阵即可
func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {return}
	r := len(board)
	c := len(board[0])
	for i:=0; i<c; i++ {   // 第一行和最后一行
		if board[0][i] == 'O' {dfs(board, 0, i)}
		if board[r-1][i] == 'O' {dfs(board, r-1, i)}
	}
	for j:=0; j<r; j++ {   // 第一列和最后一列
		if board[j][0] == 'O' {dfs(board, j, 0)}
		if board[j][c-1] == 'O' {dfs(board, j, c-1)}
	}
	for i:=0; i<r; i++ {
		for j:=0; j<c; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'W' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, x int, y int)  {
	r := len(board)
	c := len(board[0])
	if x < 0 && x >= r && y < 0 && y >= c {return}
	board[x][y] = 'W'
	if y-1 > 0 && board[x][y-1] == 'O' {dfs(board, x, y-1)}
	if x+1 < r && board[x+1][y] == 'O' {dfs(board, x+1, y)}
	if y+1 < c && board[x][y+1] == 'O' {dfs(board, x, y+1)}
	if x-1 > 0 && board[x-1][y] == 'O' {dfs(board, x-1, y)}
}

