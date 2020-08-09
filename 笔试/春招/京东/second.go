package main

import (
	"fmt"
	"io"
)

var count int

func main() {
	var n int
	var h0 int
	for {
		count = 0
		_, err := fmt.Scan(&n, &h0)
		if err == io.EOF {
			break
		}
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		dfs(arr, h0)
		fmt.Println(count % 998244353)
	}
}

func dfs(arr []int, h0 int) {
	if len(arr) == 1 {
		count++
		return
	}
	if h0 < len(arr) {
		dfs(arr[h0:], h0)
	}
	if arr[0] < len(arr) {
		dfs(arr[arr[0]:], arr[0])
	}
}
