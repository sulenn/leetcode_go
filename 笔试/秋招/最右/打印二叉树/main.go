package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		nums := make([]int, 0)
		if n == 1 {
			fmt.Println(1)
			continue
		}
		if n == 0 {
			fmt.Println(0)
			continue
		}
		nums = append(nums, 1)
		nums = append(nums, 1)
		for i := 2; i <= n; i++ {
			sum := 0
			for j := 1; j <= i; j++ {
				sum += nums[j-1] * nums[i-j]
			}
			nums = append(nums, sum)
		}
		fmt.Println(nums[n])
	}
}
