package main

//https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/

//思路，只有第一列的数向右和向下都比它大，其它的数只能判断它右边的数字比它大
func kthSmallest(matrix [][]int, k int) int {
	if k == 1 {
		return matrix[0][0]
	}
	num := 1
	rLength := len(matrix)
	cLength := len(matrix[0])
	compareArr := [][]int {{0,1}, {1,0}}
	for num != k-1 {
		minIndex := compare(matrix, compareArr)
		x := compareArr[minIndex][0]
		y := compareArr[minIndex][1]
		compareArr = append(compareArr[:minIndex], compareArr[minIndex+1:]...)
		if y == 0 && x != rLength-1 {
			compareArr = append(compareArr, []int {x+1, y})
		}
		if y != cLength-1 {
			compareArr = append(compareArr, []int {x, y+1})
		}
		num++
	}
	minIndex := compare(matrix, compareArr)    // compareArr 中最小的那个数就是第 k 小元素
	x := compareArr[minIndex][0]
	y := compareArr[minIndex][1]
	return matrix[x][y]
}

func compare(matrix [][]int,compareArr [][]int) int {   // 返回数组中最小的数的索引
	if len(compareArr) == 0 {
		return -1
	}
	min := 0
	for i:=1; i<len(compareArr); i++ {
		min_x := compareArr[min][0]
		min_y := compareArr[min][1]
		cur_x := compareArr[i][0]
		cur_y := compareArr[i][1]
		if matrix[cur_x][cur_y] < matrix[min_x][min_y] {
			min = i
		}
	}
	return min
}
