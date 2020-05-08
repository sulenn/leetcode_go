package main

//https://leetcode-cn.com/problems/maximal-square/

//思路：遍历二维矩阵中的每一个点
//每遇到一个点为1，就判断以该点为左上角起点的最大正方形是否为最大
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {return 0}
	r := len(matrix)
	c := len(matrix[0])
	count := 0
	for i:=0; i<r; i++ {
		for j:=0; j<c; j++ {
			if matrix[i][j] == '1' {  // 每遇到一个点为1，就判断以该点为左上角起点的最大正方形是否为最大
				curMax := maxArea(matrix, i, j)
				if curMax > count {
					count =  curMax
				}
			}
		}
	}
	return count
}

func maxArea(matrix [][]byte, x, y int) int {
	r := len(matrix)
	c := len(matrix[0])
	num := 1
	flag := true
	for x+num < r && y+num < c {  // 获取最大正方形的边长
		for i:=x; i<=x+num; i++ {   // 列
			if matrix[i][y+num] != '1' {
				flag = false
				break
			}
		}
		if !flag {break}
		for i:=y; i<=y+num; i++ {  // 行
			if matrix[x+num][i] != '1' {
				flag = false
				break
			}
		}
		if !flag {break}
		num++
	}
	return num * num
}