package main

import (
	"fmt"
	"io"
)

func main() {
	var N int
	for {
		_, err := fmt.Scan(&N)
		if err == io.EOF {
			break
		}
		nums1 := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&nums1[i])
		}
		nums2 := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&nums2[i])
		}
		result := 0
		for i := 0; i < N; i++ {
			result += abs(nums1[i], nums2[i])
		}
		fmt.Println(result)
	}
}

func abs(num1, num2 int) int {
	if num1-num2 < 0 {
		return num2 - num1
	}
	return num1 - num2
}
