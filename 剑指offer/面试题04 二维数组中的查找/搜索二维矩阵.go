package main

//
////https://leetcode-cn.com/problems/search-a-2d-matrix/submissions/
//
//import (
//	"fmt"
//)
//
////二分法，参考：https://leetcode-cn.com/problems/search-a-2d-matrix/solution/sou-suo-er-wei-ju-zhen-by-leetcode/
//func searchMatrix(matrix [][]int, target int) bool {
//	if len(matrix) == 0 {
//		return false
//	}
//	row := len(matrix)
//	col := len(matrix[0])
//	prev := 0
//	tail := row * col - 1
//	for tail >= prev {
//		mid := (prev + tail) / 2
//		if matrix[mid/col][mid%col] == target {
//			return true
//		}
//		if matrix[mid/col][mid%col] > target {
//			tail = mid - 1
//		} else {
//			prev = mid + 1
//		}
//	}
//	return false
//}
//
//func main() {
//	//test := make([]int,3)
//	//for i:=0;i<len(test);i++ {
//	//	test[]
//	//}
//	fmt.Println(searchMatrix([][]int{{1,   3,  5,  7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3))
//	fmt.Println(searchMatrix([][]int{{1,   3,  5,  7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13))
//	fmt.Println(searchMatrix([][]int{{1,   3}}, 3))
//}
