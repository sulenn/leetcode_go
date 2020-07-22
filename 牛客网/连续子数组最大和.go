package main

import (
	"fmt"
	"io"
	"math"
)

//https://www.nowcoder.com/practice/03d341fb6c9d42debcdd38d82a0a545c?tpId=182&&tqId=34902&rp=1&ru=/activity/oj&qru=/ta/exam-all/question-ranking

func main() {
	var a, b int

	for {
		_, err := fmt.Scan(&a)
		if err == io.EOF {
			break
		}
		var max = math.MinInt64
		var sum int
		for i := 0; i < a; i++ {
			fmt.Scan(&b)
			if sum+b > max {
				max = sum + b
			}
			if sum+b >= 0 {
				sum += b
			}
		}
		fmt.Println(max)
	}
}
