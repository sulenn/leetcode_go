package main

import (
	"fmt"
	"io"
)

func main() {
	var nums int
	var c1 int
	var c2 int
	for {
		_, err := fmt.Scan(&nums)
		if err == io.EOF {
			break
		}
		arr := make([][]int, nums)
		for i:=0; i<nums; i++ {
			fmt.Scan(&c1,&c2)
			arr[i] = []int {c1,c2}
		}
		fmt.Println(solution(arr))
	}
}

func solution(arr [][]int) int {
	money := 0
	count := 1
	visited := make([]bool,len(arr))
	for i:=0; i<len(arr);i++ {
		if arr[i][1] > 0 {
			count += arr[i][1] -1
			money += arr[i][0]
			visited[i] = true
		}
	}

	for i:=0; i<len(arr) && count > 0; i++ {
		max := -1
		for j:=0; j<len(arr); j++ {
			if visited[j] {
				continue
			}
			if !visited[j] && max == -1 {
				max = j
			} else if arr[j][0] > arr[max][0] && !visited[j] {
				max = j
			}
		}
		if max == -1 {return money}
		count--
		money += arr[max][0]
		visited[max] = true
	}
	return money
}