package main

//https://leetcode-cn.com/problems/game-of-life/

func gameOfLife(board [][]int)  {
	if len(board) == 0 {
		return
	}
	newBoard := make([][]int, len(board))  // 可以不用额外数组，用一个复合状态表示。比如用2表示从0变为1，用3表示从1变为0
	for i:=0; i<len(board); i++ {
		newBoard[i] = make([]int, len(board[0]))
		copy(newBoard[i],board[i])
	}
	location := [][]int {{-1,0},{-1,1},{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1}}  // 8个方向
	for i:=0; i<len(newBoard); i++ {
		for j:=0; j<len(newBoard[0]); j++ {
			oneNums := 0
			for v:=0; v<8;v++ {
				if i+location[v][0] >=0 && i+location[v][0] <len(board) && j+location[v][1] >=0 && j+location[v][1] <len(newBoard[0]) && newBoard[i+location[v][0]][j+location[v][1]] == 1 {
					oneNums += 1
				}
			}
			if oneNums < 2 || oneNums > 3 {
				board[i][j] = 0
			} else if oneNums == 3 {
				board[i][j] = 1
			}
		}
	}
}
