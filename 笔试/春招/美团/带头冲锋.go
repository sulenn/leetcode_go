package main

import (
	"fmt"
	"io"
)

func main() {
	var num int
	var location int
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
		dic := make(map[int]int)
		for i := 0; i < num; i++ {
			fmt.Scan(&location)
			dic[location] = i
		}
		count := 0
		arr := make([]int, num)
		for i := 0; i < num; i++ {
			fmt.Scan(&location)
			arr[i] = location
		}
		for i := 0; i < num; i++ {
			for j := i + 1; j < num; j++ {
				if dic[arr[i]] > dic[arr[j]] {
					count++
					break
				}
			}
		}
		fmt.Println(count)
	}
}
