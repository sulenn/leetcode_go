package main

import (
	"fmt"
)

//https://www.nowcoder.com/test/question/fe298c55694f4ed39e256170ff2c205f?pid=1088888&tid=31498494

func main() {
	for {
		var a int
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		fmt.Println(a / 2)
	}
}
