package main

//https://leetcode-cn.com/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	lineArrs := make([][]int, 0)
	columnArrs := make([][]int, 0)

	length := len(board)
	width := len(board[0])
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {

		}
	}
}

func fill(board [][]byte, x, y int) bool {
	if board[x][y] != '.' {
		if x < 8 {
			return fill(board, x+1, y)
		}

	}
}
