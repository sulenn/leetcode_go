package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	var N int
	var value int
	for {
		_, err := fmt.Scan(&N)
		if err == io.EOF {
			break
		}
		arr := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&value)
			arr[i] = value
		}
		sum := 0
		sort.Ints(arr)
		count := 1
		for i := N - 1; i >= 0; i-- {
			if count%3 != 0 {
				sum += arr[i]
			}
			count++
		}
		fmt.Println(sum)
	}
}
