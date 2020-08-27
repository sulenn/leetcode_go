package main

import (
	"fmt"
	"io"
)

var max = 0

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		arr := make([][]int, n)
		for i := 0; i < n; i++ {
			arr[i] = make([]int, 2*n-1)
		}
		for i := 0; i < n; i++ {
			start := n - i - 1
			for start < n+i {
				fmt.Scan(&arr[i][start])
				start++
			}
		}
		max = 0
		dfs(0, n-1, 0, arr, arr[0][n-1])
		fmt.Println(max)
	}
}

func dfs(x, y, d int, arr [][]int, curValue int) {
	if d >= len(arr)-1 {
		return
	}
	left := len(arr) - (d + 1) - 1
	right := len(arr) + (d + 1) - 1
	if curValue+arr[x+1][y] > max {
		max = curValue + arr[x+1][y]
	}
	dfs(x+1, y, d+1, arr, curValue+arr[x+1][y])
	if y-1 >= left {
		if curValue+arr[x+1][y-1] > max {
			max = curValue + arr[x+1][y-1]
		}
		dfs(x+1, y-1, d+1, arr, curValue+arr[x+1][y-1])
	}
	if y+1 <= right {
		if curValue+arr[x+1][y+1] > max {
			max = curValue + arr[x+1][y+1]
		}
		dfs(x+1, y+1, d+1, arr, curValue+arr[x+1][y+1])
	}
}
