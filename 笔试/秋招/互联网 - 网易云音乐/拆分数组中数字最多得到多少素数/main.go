package main

import (
	"fmt"
	"io"
	"math/big"
)

// 大概率是 int 溢出了，导致出错。应该使用 int64

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		var num int
		var result = big.NewInt(0)
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			num = num >> 1
			temp := big.NewInt(int64(num))
			result.Add(result, temp)
		}
		fmt.Println(result.String())
	}
}
