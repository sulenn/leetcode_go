package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	var m int
	for {
		_, err := fmt.Scan(&n, &m)
		if err ==io.EOF {break}
		arr := make([]int,n)
		for i:=0; i<n; i++ {
			fmt.Scan(&arr[i])
		}
		if n == 1 {
			fmt.Println(m)
			continue
		}
		//if arr[1] > arr[0] {
		//	arr[1], arr[0] = arr[0], arr[1]
		//}
		count := 0
		num := m
		max := arr[0]
		for i:=1; i<n && num > 0; i++ {
			if max > arr[i] {
				num--
			} else {
					max = arr[i]
					num = m-1
			}
			count++
		}
		fmt.Println(count+num)
	}
}
