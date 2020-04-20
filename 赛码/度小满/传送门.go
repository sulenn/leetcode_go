package main

import (
	"fmt"
	"io"
)

func main() {
	var N int
	var A int
	var B int
	var C int
	for {
		_, err := fmt.Scan(&N, &A, &B, &C)
		if err == io.EOF {break}
		arr := make([]int, N)
		for i:=0;i<N;i++ {
			fmt.Scan(&arr[i])
		}
		fmt.Println(solution1(arr, A, B, C))
	}
}

func solution1(arr []int, A, B, C int) int {
	if arr[0] == len(arr) {return A}
	if arr[0] > len(arr) {return (arr[0]-len(arr))*C + A}
	if arr[0] < len(arr) {return (len(arr)-arr[0])*B + A}
	return 0
}
