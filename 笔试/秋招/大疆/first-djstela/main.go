package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var N, P, X int
	for {
		_, err := fmt.Scan(&N, &P)
		if err == io.EOF {
			break
		}
		graph := make([][]int, N)
		for i := 0; i < N; i++ {
			graph[i] = make([]int, N)
			for j := 0; j < N; j++ {
				graph[i][j] = math.MaxInt64
			}
		}
		var num1, num2, value int
		for i := 0; i < P; i++ {
			fmt.Scan(&num1, &num2, &value)
			graph[num1][num2] = value
		}
		fmt.Scan(&X)
		if X == 0 || X == N {
			fmt.Println(0)
			continue
		}
		fmt.Println(findMinPath(graph, X))
	}
}

func findMinPath(graph [][]int, target int) int {
	var curMinPath int
	var curMinPathIndex int
	var isGetPathArr = make([]bool, len(graph))
	var minPathArr = make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		minPathArr[i] = graph[0][i]
	}
	isGetPathArr[0] = true
	minPathArr[0] = 0
	for i := 1; i < len(graph); i++ {
		curMinPath = math.MaxInt64
		for j := 0; j < len(graph); j++ {
			if !isGetPathArr[j] && minPathArr[j] < curMinPath {
				curMinPath = minPathArr[i]
				curMinPathIndex = j
			}
		}
		if curMinPath == math.MaxInt64 {
			break
		}
		isGetPathArr[curMinPathIndex] = true
		for j := 0; j < len(graph); j++ {
			if !isGetPathArr[j] && curMinPath+graph[curMinPathIndex][j] < minPathArr[j] {
				minPathArr[j] = curMinPath + graph[curMinPathIndex][j]
			}
		}
	}
	return minPathArr[target]
}
