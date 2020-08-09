package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var n int
	var m int
	var a int
	var b int
	for {
		_, err := fmt.Scan(&n, &m, &a, &b)
		if err == io.EOF {
			break
		}
		if a > n || b > m {
			fmt.Println(0)
			continue
		}
		arr1 := make([][]int, n)
		for i := 0; i < n; i++ {
			arr1[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				arr1[i][j] = (i + 1) * (j + 1) % 10
			}
		}
		fmt.Println(solution(arr1, a, b))
	}
}

func solution(arr1 [][]int, a, b int) int {
	sum := 0
	for i := 0; i+a <= len(arr1); i++ {
		for j := 0; j+b <= len(arr1[0]); j++ {
			sum += mul(arr1, i, i+a, j, j+b)
		}
	}
	return sum
}

func mul(arr [][]int, r1, r2, c1, c2 int) int {
	max := math.MinInt64
	for i := r1; i < r2; i++ {
		for j := c1; j < c2; j++ {
			if arr[i][j] > max {
				max = arr[i][j]
			}
			if max == 9 {
				return 9
			}
		}
	}
	return max
}
