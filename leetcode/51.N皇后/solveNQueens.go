package main

//https://leetcode-cn.com/problems/n-queens/

func solveNQueens(n int) [][]string {
	if n < 1 {
		return [][]string{}
	}
	arr := make([][]int, n) // 0 为空白，1为皇后，大于等于2为冲突区
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	return back(arr, 0, n)
}

func back(arr [][]int, curRow int, n int) [][]string {
	if curRow == n { // 终止条件
		strs := []string{}
		for i := 0; i < n; i++ { // 组织二维切片为字符串
			str := ""
			for j := 0; j < n; j++ {
				if arr[i][j] == 1 {
					str += "Q"
				} else {
					str += "."
				}
			}
			strs = append(strs, str)
		}
		return [][]string{strs}
	}
	result := [][]string{}
	for i := 0; i < n; i++ {
		if arr[curRow][i] == 0 {
			addCondition(arr, curRow, i, n) // 添加皇后点约束
			result = append(result, back(arr, curRow+1, n)...)
			deleteCondition(arr, curRow, i, n) // 回溯，删除皇后点约束
		}
	}
	return result
}

// 添加皇后点产生的约束
func addCondition(arr [][]int, x int, y int, n int) { // 之所以2+x，目的是区分0和1，以及回溯时删除攻击区域
	for i := 0; i < n; i++ { // 将皇后点同行和同列设置为攻击约束区
		arr[x][i] += 2 + x
		arr[i][y] += 2 + x
	}
	for i := 1; i < n; i++ { // 将皇后点对角线设置为攻击约束区
		if x+i < n && y+i < n { // 右上
			arr[x+i][y+i] += 2 + x
		}
		if x-i >= 0 && y-i >= 0 { // 左下
			arr[x-i][y-i] += 2 + x
		}
		if x+i < n && y-i >= 0 { // 左上
			arr[x+i][y-i] += 2 + x
		}
		if x-i >= 0 && y+i < n { // 右下
			arr[x-i][y+i] += 2 + x
		}
	}
	arr[x][y] = 1 // 设置皇后点
}

// 删除皇后点产生的约束
func deleteCondition(arr [][]int, x int, y int, n int) {
	for i := 0; i < n; i++ { // 将皇后点同行和同列设置为攻击约束区
		arr[x][i] -= 2 + x
		arr[i][y] -= 2 + x
	}
	for i := 1; i < n; i++ { // 将皇后点对角线设置为攻击约束区
		if x+i < n && y+i < n { // 右上
			arr[x+i][y+i] -= 2 + x
		}
		if x-i >= 0 && y-i >= 0 { // 左下
			arr[x-i][y-i] -= 2 + x
		}
		if x+i < n && y-i >= 0 { // 左上
			arr[x+i][y-i] -= 2 + x
		}
		if x-i >= 0 && y+i < n { // 右下
			arr[x-i][y+i] -= 2 + x
		}
	}
	arr[x][y] = 0 // 设置皇后点
}
