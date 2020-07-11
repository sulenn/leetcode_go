package main

import "fmt"

//https://leetcode-cn.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	matrix := make([][]int, m) //二维切片
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := range matrix { //第一列全赋值为 0
		matrix[i][0] = 1
	}
	for i := range matrix[0] { //第一行全赋值为 0
		matrix[0][i] = 1
	}
	for i := 1; i < m; i++ { //动态规划，matrix[m][n] = matrix[m-1][n] + matrix[m][n-1]
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[m-1][n-1]
}

//递归的方法会超出时间限制，在 m = 25, n = 9
//func recursive(m int, n int, x int, y int) int {
//	if x == m || y == n{
//		return 1
//	}
//	up := recursive(m, n, x, y + 1)
//	right := recursive(m, n, x + 1, y)
//	return up + right
//}

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(25, 9))
}
