package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF{break}
		arr := make([][]int,n)
		for i:=0; i<n; i++ {
			var num int
			nums := []int {}
			for j:=0; j<=i; j++ {
				fmt.Scan(&num)
				nums = append(nums, num)
			}
			arr[i] = nums
		}
		fmt.Println(solution1(arr))
	}
}

func solution1(triangle [][]int) int {
	tLen:=len(triangle)
	if tLen<=0{
		return 0
	}

	for i:=tLen-2;i>=0;i--{
		for j:=len(triangle[i])-1;j>=0;j--{
			triangle[i][j]=triangle[i][j]+max(triangle[i+1][j],triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func max(a,b int) int{
	if a>b{
		return a
	}

	return b
}

