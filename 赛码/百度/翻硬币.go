package main

import (
	"fmt"
	"io"
)

func main() {
	var N int
	var x int
	var y int
	var coin int
	for {
		_, err := fmt.Scan(&N, &x, &y)
		if err == io.EOF {
			break
		}
		coins := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&coin)
			coins[i] = coin
		}
		fmt.Println(solution1(x, y, coins))
	}
}

func solution1(x int, y int, coins []int) int {
	sum := 0
	if x <= y { // 当 x 小于 y时
		start := 0
		for i := 0; i < len(coins); i++ {
			if coins[i] == 0 {
				continue
			}
			if coins[i] != 0 && coins[start] == 0 {
				sum += x
			}
			start = i + 1
		}
		if start != len(coins) {
			sum += x
		}
	} else {
		start := 0
		for i := 0; i < len(coins); i++ {
			if coins[i] == 0 {
				continue
			}
			if coins[i] != 0 && coins[start] == 0 {
				sum += y
			}
			start = i + 1
		}
		if start != len(coins) {
			sum += y
		}
	}
	return sum
}
