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
		result := 0
		for i := 5; i <= n; i += 5 {
			temp := i
			for temp%5 == 0 {
				result++
				temp /= 5
			}
		}
		fmt.Println(result)
	}
}
