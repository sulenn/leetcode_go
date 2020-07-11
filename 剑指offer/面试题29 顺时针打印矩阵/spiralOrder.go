package main

//https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	rowPrev := 0
	rowTail := len(matrix) - 1
	colPrev := 0
	colTail := len(matrix[0]) - 1
	result := make([]int, len(matrix)*len(matrix[0]))
	count := 0
	for rowPrev < rowTail && colPrev < colTail { // 顺时针遍历
		for i := colPrev; i <= colTail; i++ { // 上
			result[count] = matrix[rowPrev][i]
			count++
		}
		rowPrev++
		for i := rowPrev; i <= rowTail; i++ { // 右
			result[count] = matrix[i][colTail]
			count++
		}
		colTail--
		for i := colTail; i >= colPrev; i-- { // 下
			result[count] = matrix[rowTail][i]
			count++
		}
		rowTail--
		for i := rowTail; i >= rowPrev; i-- { // 左
			result[count] = matrix[i][colPrev]
			count++
		}
		colPrev++
	}
	if rowPrev == rowTail { // 剩下一行
		for i := colPrev; i <= colTail; i++ {
			result[count] = matrix[rowPrev][i]
			count++
		}
	} else if colPrev == colTail { // 剩下一列
		for i := rowPrev; i <= rowTail; i++ {
			result[count] = matrix[i][colPrev]
			count++
		}
	}
	return result
}
