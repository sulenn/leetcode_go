package main

import (
	"fmt"
	"io"
)
var count int

func main() {
	var N int
	var M int
	var value int
	for {
		count = 0
		_, err := fmt.Scan(&N,&M)
		if err == io.EOF {break}
		arr := make([]int, N)
		for i:=0; i<N; i++ {
			fmt.Scan(&value)
			arr[i] = value
		}
		for i:=0; i<N; i++ {
			solution(arr[i+1:],M,arr[i])
		}
		fmt.Println(count)
	}
}

func solution(arr []int, M int, sum int)  {
	if sum % M == 0 {
		count++
	}
	if len(arr) == 0 {return}
	solution(arr[1:], M, sum+arr[0])
}

