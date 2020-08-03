package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var lunchCount int
	var dinnerCount int
	var deliciousTarget int
	for {
		_, err := fmt.Scan(&lunchCount, &dinnerCount, &deliciousTarget)
		if err == io.EOF {
			break
		}
		lunchNums := make([][2]int, lunchCount)
		dinnerNums := make([][2]int, dinnerCount)
		min := math.MaxInt64
		for i := 0; i < lunchCount; i++ {
			fmt.Scan(&lunchNums[i][0], &lunchNums[i][1])
			if lunchNums[i][1] >= deliciousTarget && lunchNums[i][0] < min {
				min = lunchNums[i][0]
			}
		}
		for i := 0; i < dinnerCount; i++ {
			fmt.Scan(&dinnerNums[i][0], &dinnerNums[i][1])
			if dinnerNums[i][1] >= deliciousTarget && dinnerNums[i][0] < min {
				min = dinnerNums[i][0]
			}
		}
		if min != math.MaxInt64 {
			fmt.Println(min)
			continue
		}
		if deliciousTarget == 0 {
			fmt.Println(0)
			continue
		}
		for i := 0; i < lunchCount; i++ {
			for j := 0; j < dinnerCount; j++ {
				if lunchNums[i][1]+dinnerNums[j][1] >= deliciousTarget && lunchNums[i][0]+dinnerNums[j][0] < min {
					min = lunchNums[i][0] + dinnerNums[j][0]
				}
			}
		}
		if min == math.MaxInt64 {
			fmt.Println(-1)
		} else {
			fmt.Println(min)
		}
	}
}
