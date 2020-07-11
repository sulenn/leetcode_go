package main

//https://leetcode-cn.com/problems/available-captures-for-rook/

//找到白色小车后，从上下左右四个方向匹配卒即可
func numRookCaptures(board [][]byte) int {
	R := make([]int, 2) // 白色的车
	flag := false
	for i := 0; i < len(board); i++ { // 找白色车位置
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				R[0] = i
				R[1] = j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	result := 0
	for i := 1; R[1]-i >= 0; i++ { // 左
		if board[R[0]][R[1]-i] == 'B' {
			break
		}
		if board[R[0]][R[1]-i] == 'p' {
			result++
			break
		}
	}
	for i := 1; R[0]+i < 8; i++ { // 上
		if board[R[0]+i][R[1]] == 'B' {
			break
		}
		if board[R[0]+i][R[1]] == 'p' {
			result++
			break
		}
	}
	for i := 1; R[1]+i < len(board[0]); i++ { // 右
		if board[R[0]][R[1]+i] == 'B' {
			break
		}
		if board[R[0]][R[1]+i] == 'p' {
			result++
			break
		}
	}
	for i := 1; R[0]-i >= 0; i++ { // 下
		if board[R[0]-i][R[1]] == 'B' {
			break
		}
		if board[R[0]-i][R[1]] == 'p' {
			result++
			break
		}
	}
	return result
}
