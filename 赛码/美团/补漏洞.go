package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	var k int
	for {
		_, err := fmt.Scan(&n, &k)
		if err == io.EOF{break}
		//min := n
		pre := 0
		tail := n

		for pre != tail {
			mid := pre + (tail-pre)/2
			if solution3(mid,n,k) {
				tail = mid
			} else {
				pre = mid+1
			}
		}
		fmt.Println(pre)
		//count := 1
		//for !solution3(count, n, k) {
		//	count++
		//}
		//fmt.Println(count)
	}
}

func solution3(num, n,k int) bool {
	sum := 1
	for n > 0 && num/sum >= 1 {
		n -= num/sum
		sum *= k
	}
	if n <= 0 {return true}
	return false
}
