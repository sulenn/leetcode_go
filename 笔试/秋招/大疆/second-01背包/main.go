package main

import (
	"fmt"
	"io"
)

func main() {
	var N, X int
	for {
		_, err := fmt.Scan(&N, &X)
		if err == io.EOF {
			break
		}
		numsArr := make([][2]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&numsArr[i][0], &numsArr[i][1])
		}
		result := make([][]int, N+1)
		for i := 0; i < N+1; i++ {
			result[i] = make([]int, X+1)
		}
		for i := numsArr[0][1]; i < X+1; i++ {
			result[1][i] = numsArr[0][0]
		}
		for i := 2; i < N+1; i++ {
			for j := numsArr[i-1][1]; j < X+1; j++ {
				if result[i-1][j] > result[i-1][j-numsArr[i-1][1]]+numsArr[i-1][0] {
					result[i][j] = result[i-1][j]
				} else {
					result[i][j] = result[i-1][j-numsArr[i-1][1]] + numsArr[i-1][0]
				}
			}
		}
		fmt.Println(result[N][X])
	}
}
