package main

import (
	"fmt"
	"io"
)

func main() {
	var s string
	var Q int
	for {
		_, err := fmt.Scan(&s)
		if err == io.EOF {
			break
		}
		fmt.Scan(&Q)
		var num1, num2 int
		for i := 0; i < Q; i++ {
			fmt.Scan(&num1, &num2)
			fmt.Println(minCut(s[num1-1 : num2-1]))
		}
	}
}

func minCut(str string) int {
	length := len(str)
	f := make([]int, length+1)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	for i := 0; i <= length; i++ {
		f[i] = length - 1 - i
	}
	for i := length - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i; j < length; j++ {
			if str[i] == str[j] && ((j-i < 2) || dp[i+1][j-1] == 1) {
				dp[i][j] = 1
				if f[i] > f[j+1]+1 {
					f[i] = f[j+1] + 1
				}
			}
		}
	}
	return f[0]
}
