package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	var num int
	var time string
	var times [][]int
	for {
		times = [][]int{}
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
		for i := 0; i < num; i++ {
			fmt.Scan(&time)
			arr := strings.Split(time, ",")
			times = append(times, []int{})
			for j := 0; j < 2; j++ {
				t, _ := strconv.Atoi(arr[j])
				times[i] = append(times[i], t)
			}
		}
		fmt.Println(solution(times))
	}
}

func solution(times [][]int) int {
	visited := make([]bool, len(times))
	result := 0
	for i := 0; i < len(times); i++ {
		if visited[i] {
			continue
		}
		cur := times[i]
		for j := i; j < len(times); j++ {
			if cur[1] <= times[j][0] && visited[j] == false {
				cur = times[j]
				visited[j] = true
			}
		}
		result++
	}
	return result
}
