package main

import (
	"fmt"
	"io"
)

func main() {
	var n, m int64
	for {
		_, err := fmt.Scan(&n, &m)
		if err == io.EOF {
			break
		}
		nums := make([]int64, m)
		var i int64 = 0
		for ; i < m; i++ {
			fmt.Scan(&nums[i])
		}
		result := 0
		for i = 1; i <= n; i++ {
			for j := 0; j < len(nums); j++ {
				if i%nums[j] == 0 {
					result++
					break
				}
			}
		}
		fmt.Println(result)
	}
}
