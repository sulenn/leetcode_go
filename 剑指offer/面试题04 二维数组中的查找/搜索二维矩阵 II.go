package main

import "fmt"

//https://leetcode-cn.com/problems/search-a-2d-matrix-ii/

//选左上角，往右走和往下走都增大，不能选
//
//选右下角，往上走和往左走都减小，不能选
//
//选左下角，往右走增大，往上走减小，可选
//
//选右上角，往下走增大，往左走减小，可选
//参考方法4：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/solution/sou-suo-er-wei-ju-zhen-ii-by-leetcode-2/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	leftDownRow := row - 1
	leftDownCol := 0
	for leftDownRow >= 0 && leftDownCol < col {
		if matrix[leftDownRow][leftDownCol] == target {
			return true
		}
		if matrix[leftDownRow][leftDownCol] > target {
			leftDownRow--
		} else {
			leftDownCol++
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}
