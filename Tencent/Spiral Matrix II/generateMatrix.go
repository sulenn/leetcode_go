package main

import "fmt"

//https://leetcode-cn.com/problems/spiral-matrix-ii/

func generateMatrix(n int) [][]int {
	upRow := 0
	downRow := n - 1
	leftLine := 0
	rightLine := n - 1
	count := 1                 //计数器
	result := make([][]int, n) //申明 n 维数组
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for downRow >= upRow {
		for i := leftLine; i <= rightLine; i++ { // 上
			result[upRow][i] = count
			count++
		}
		upRow++
		for i := upRow; i <= downRow; i++ { // 右
			result[i][rightLine] = count
			count++
		}
		rightLine--
		for i := rightLine; i >= leftLine; i-- { // 下
			result[downRow][i] = count
			count++
		}
		downRow--
		for i := downRow; i >= upRow; i-- { // 左
			result[i][leftLine] = count
			count++
		}
		leftLine++
	}
	return result
}

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(1))
}
