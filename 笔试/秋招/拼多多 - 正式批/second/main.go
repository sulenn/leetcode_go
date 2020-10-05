package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

var max int
var result []int

func main() {
	var N, M int
	for {
		max = math.MaxInt64
		result = []int{}
		_, err := fmt.Scan(&N, &M)
		if err == io.EOF {
			break
		}
		arr := make([][]byte, N)
		for i := 0; i < N; i++ {
			arr[i] = make([]byte, M)
		}
		var str string
		for i := 0; i < N; i++ {
			fmt.Scan(&str)
			arr[i] = []byte(str)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if arr[i][j] == 'X' {
					dfs(arr, i, j, i, j, N, M, -1)
				}
			}
		}
		if len(result) == 0 {
			fmt.Println(0)
			break
		}
		fmt.Println(max)
		res := ""
		for i := 0; i < len(result); i++ {
			res += strconv.Itoa(result[i]) + " "
		}
		fmt.Println(strings.TrimSpace(res))
	}
}

func dfs(arr [][]byte, originX, originY, curX, curY, N, M, length int) {
	length++
	if arr[curX][curY] == 'T' {
		if length < max {
			result = []int{}
			max = length
		}
		if len(result) > 0 && result[len(result)-2] == originX && result[len(result)-1] == originY { // 同一个传送点不能有多条路径
			return
		}
		if length == max {
			result = append(result, originX, originY)
		}
		return
	}
	if length > max {
		return
	}
	if arr[curX][curY] == '1' {
		return
	}
	temp := arr[curX][curY]
	arr[curX][curY] = '1'
	if curX+1 < N {
		dfs(arr, originX, originY, curX+1, curY, N, M, length)
	}
	if curY+1 < M {
		dfs(arr, originX, originY, curX, curY+1, N, M, length)
	}
	if curX-1 >= 0 {
		dfs(arr, originX, originY, curX-1, curY, N, M, length)
	}
	if curY-1 >= 0 {
		dfs(arr, originX, originY, curX, curY-1, N, M, length)
	}
	arr[curX][curY] = temp
}
