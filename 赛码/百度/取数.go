package main

import (
	"fmt"
	"io"
)

var max int

func main() {
	var n int
	var m int
	var num int
	for {
		max = 0
		_, err := fmt.Scan(&n,&m)
		if err == io.EOF {
			break
		}
		numArr1 := make([]int, n)
		for i:=0; i<n; i++ {
			fmt.Scan(&num)
			numArr1[i] = num
		}
		numArr2 := make([]int, n)
		for i:=0; i<n; i++ {
			fmt.Scan(&num)
			numArr2[i] = num
		}
		visited := make([]bool, n)
		solution(visited,m,numArr1,numArr2,0,0)
		fmt.Println(max)
	}
}

func solution(visited []bool, m int, numArr1 []int, numArr2 []int, sum int, minus int) {
	if m == 0 {
		if sum > max {
			max = sum
		}
		return
	}
	for i:=0; i<len(numArr1); i++ {
		if visited[i] {continue}
		visited[i] = true
		solution(visited,m-1,numArr1,numArr2,sum+numArr1[i]-minus, minus+numArr2[i])
		visited[i] =false
	}
}
