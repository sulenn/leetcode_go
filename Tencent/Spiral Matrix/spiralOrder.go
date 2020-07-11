package main

import "fmt"

//https://leetcode-cn.com/problems/spiral-matrix/

//一个正方形边框一个正方形边框的遍历，像剥洋葱一样，不过需要对一行或者一列的情况分开讨论
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	upRow := 0                      //矩阵上边界
	downRow := len(matrix) - 1      //矩阵下边界
	leftLine := 0                   //矩阵左边界
	rightLine := len(matrix[0]) - 1 //矩阵右边界
	result := []int{}
	for downRow-upRow > 0 && rightLine-leftLine > 0 { //剥完，或者剩下一行、或一列。之所以剩下一行或一列是为了防止重复遍历
		for i := leftLine; i <= rightLine; i++ { //矩阵上边界
			result = append(result, matrix[upRow][i])
		}
		upRow++
		for i := upRow; i <= downRow; i++ { //矩阵下边界
			result = append(result, matrix[i][rightLine])
		}
		rightLine--
		for i := rightLine; i >= leftLine; i-- { //矩阵左边界
			result = append(result, matrix[downRow][i])
		}
		downRow--
		for i := downRow; i >= upRow; i-- { //矩阵右边界
			result = append(result, matrix[i][leftLine])
		}
		leftLine++
	}
	if downRow-upRow == 0 { //剩一行
		for i := leftLine; i <= rightLine; i++ {
			result = append(result, matrix[upRow][i])
		}
	} else if rightLine-leftLine == 0 { //剩一列
		for i := upRow; i <= downRow; i++ {
			result = append(result, matrix[i][rightLine])
		}
	}
	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3, 4}}))
}
