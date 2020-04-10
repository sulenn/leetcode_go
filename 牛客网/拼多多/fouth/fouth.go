package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var N int
	var K int
	var value int
	for {
		_, err := fmt.Scan(&N, &K)
		if err == io.EOF {break}
		arr := make([]int, N)
		for i:=0; i<N; i++ {
			fmt.Scan(&value)
			arr[i] = value
		}
		dic := make(map[int]int, 0)
		for i:=0 ; i<N; i++ {
			dic[arr[i]]++
		}
		max := math.MinInt64
		for _,v := range dic {
			if v > max {
				max = v
			}
		}
		fmt.Println(max)
	}
}
