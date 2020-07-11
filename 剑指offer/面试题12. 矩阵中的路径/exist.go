package main

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if recusive(board, i, j, word) {
					return true
				}
			}
		}
	}
	return false
}

func recusive(board [][]byte, x int, y int, word string) bool {
	row := len(board) - 1
	col := len(board[0]) - 1
	if x < 0 || x > row || y < 0 || y > col || len(word) == 0 {
		return false
	}
	if board[x][y] == word[0] {
		if len(word) == 1 {
			return true
		}
		tempValue := board[x][y]
		board[x][y] = '-'
		if y-1 >= 0 && board[x][y-1] != '-' { // 左
			if recusive(board, x, y-1, word[1:]) {
				return true
			}
		}
		if y+1 <= col && board[x][y+1] != '-' { // 右
			if recusive(board, x, y+1, word[1:]) {
				return true
			}
		}
		if x-1 >= 0 && board[x-1][y] != '-' { // 上
			if recusive(board, x-1, y, word[1:]) {
				return true
			}
		}
		if x+1 <= row && board[x+1][y] != '-' { // 下
			if recusive(board, x+1, y, word[1:]) {
				return true
			}
		}
		board[x][y] = tempValue
	}
	return false
}

//func main()  {
//	temp := [][]int {{1,2}}
//	exist(temp, "sd")
//	fmt.Println(temp)
//}
