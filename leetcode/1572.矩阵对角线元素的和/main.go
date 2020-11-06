package main

//https://leetcode-cn.com/problems/matrix-diagonal-sum/

func diagonalSum(mat [][]int) int {
	result := 0
	length := len(mat)
	for i := 0; i < length; i++ {
		first := i
		second := length - i - 1
		if first == second {
			result += mat[i][first]
		} else {
			result += mat[i][first] + mat[i][second]
		}
	}
	return result
}
