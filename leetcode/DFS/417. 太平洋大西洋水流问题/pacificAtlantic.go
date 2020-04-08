package main

import (
	"math"
)

//https://leetcode-cn.com/problems/pacific-atlantic-water-flow/

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {return [][]int{}}
	result := make([][]int, 0)
	newMatrix := make([][]int, len(matrix))
	for i:=0; i<len(matrix); i++ {
		newMatrix[i] = make([]int, len(matrix[0]))
	}
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			for v:=0; v<len(matrix); v++ {
				copy(newMatrix[v], matrix[v])
			}
			if dfs(newMatrix, i, j) == 3 {
				result = append(result, []int {i, j})
			}
		}
	}
	return result
}

func dfs(matrix [][]int, i int, j int) int {   // 0 为空，1为太平洋，2为大西洋，3为两者皆可
	if (i == 0 && j == len(matrix[0])-1) || (i == len(matrix)-1 && j == 0) {return 3}
	if i == len(matrix)-1 || j == len(matrix[0])-1 {return 2}
	if j == 0 || i == 0 {return 1}
	location := [][]int {{0,-1},{1,0},{0,1},{-1,0}}   // 左、上、右、下
	state := 0
	curValue := matrix[i][j]
	matrix[i][j] = math.MaxInt64
	for v:=0; v<4; v++ {
		r := i+location[v][0]
		c := j+location[v][1]
		if (r>=0 && r<len(matrix)) && (c>=0 && c<len(matrix[0])) && curValue >= matrix[r][c] {
			back := dfs(matrix, r, c)
			if back == 3 {return 3}
			if (back==1 && state==2) || (back==2 && state==1) {return 3}
			state = back
		}
	}
	return state
}

//func main() {
//	test := [][]int {{},{},{},{}}
//	fmt.Println(len(test))
//	fmt.Println(len(test[0]))
//}