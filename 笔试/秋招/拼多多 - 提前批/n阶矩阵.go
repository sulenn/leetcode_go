package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		mid := float64(n-1) / 2
		var str string
		for x := 0; x < n; x++ {
			str = ""
			for y := 0; y < n; y++ {
				if x+y == n-1 {
					str += "0 "
					continue
				}
				if x == y {
					str += "0 "
					continue
				}
				if float64(x) == mid || float64(y) == mid {
					str += "0 "
					continue
				}
				if float64(x) > mid && float64(y) > mid && x < y {
					str += "7 "
				}
				if float64(x) < mid && float64(y) > mid && x+y > n-1 {
					str += "8 "
				}
				if float64(x) < mid && float64(y) > mid && x+y < n-1 {
					str += "1 "
				}
				if float64(x) < mid && float64(y) < mid && x < y {
					str += "2 "
				}
				if float64(x) < mid && float64(y) < mid && x > y {
					str += "3 "
				}
				if float64(x) > mid && float64(y) < mid && x+y < n-1 {
					str += "4 "
				}
				if float64(x) > mid && float64(y) < mid && x+y > n-1 {
					str += "5 "
				}
				if float64(x) > mid && float64(y) > mid && x > y {
					str += "6 "
				}
			}
			str = strings.TrimRight(str, " ")
			fmt.Println(str)
		}
	}
}
