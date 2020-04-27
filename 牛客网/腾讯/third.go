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
		arr := make([][2]int, n)
		for i:=0; i<n; i++ {
			fmt.Scan(&arr[i][0])
		}
		for i:=0; i<n; i++ {
			fmt.Scan(&arr[i][1])
		}
		fmt.Println(-1)
	}
}
