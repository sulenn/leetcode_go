package main

//https://leetcode-cn.com/problems/word-search/

func exist(board [][]byte, word string) bool {
	length := len(board)
	width := len(board[0])
	flag := make([][]bool, length)
	for v := 0; v < length; v++ {
		flag[v] = make([]bool, width)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == word[0] && backtracking(board, word[1:], i, j, flag) {
				return true
			}
		}
	}
	return false
}

func backtracking(board [][]byte, word string, x, y int, flag [][]bool) bool {
	if len(word) == 0 {
		return true
	}
	flag[x][y] = true
	height := len(board)
	width := len(board[0])
	if x+1 < height && !flag[x+1][y] && board[x+1][y] == word[0] && backtracking(board, word[1:], x+1, y, flag) {
		return true
	}
	if y-1 >= 0 && !flag[x][y-1] && board[x][y-1] == word[0] && backtracking(board, word[1:], x, y-1, flag) {
		return true
	}
	if x-1 >= 0 && !flag[x-1][y] && board[x-1][y] == word[0] && backtracking(board, word[1:], x-1, y, flag) {
		return true
	}
	if y+1 < width && !flag[x][y+1] && board[x][y+1] == word[0] && backtracking(board, word[1:], x, y+1, flag) {
		return true
	}
	flag[x][y] = false
	return false
}
