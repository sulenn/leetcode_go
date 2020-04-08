package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	var m int
	var a int
	var b int
	for {
		_, err := fmt.Scan(&n,&m,&a,&b)
		if err == io.EOF {
			break
		}
		if n == 0 {
			fmt.Println(0)
			break
		}
		fmt.Println(solution(n,m,a,b))
	}
}

func solution(n,m,a,b int) int {
	count := 0
	if n >= m {
		if a > b {
			count = (n-m)*b
		} else {
			count = (n-m)*a
		}
	} else {
		count = (m % n)*b
		num := 0
		for m%n != 0 {
			num++
			n--
		}
		if num*a < count{
			count = num*a
		}
	}
	return count
}
