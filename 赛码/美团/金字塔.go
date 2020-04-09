package main

import (
	"fmt"
	"io"
)

func main() {
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {break}
		if num == 0 {
			fmt.Println(1)
			continue
		}
		if num == 1 {
			fmt.Println(0)
			continue
		}
		if num == 2 {
			fmt.Println(3)
			continue
		}
		dp1 := 3
		dp2 := 6
		for i:=4; i<num; i++ {
			dp2,dp1 = (dp1*dp2) %1000000007, dp2
		}
		fmt.Println(dp2)
	}
}


