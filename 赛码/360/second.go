package main

import (
	"fmt"
	"io"
)

var percent float64

func main() {
	var n int
	var a int
	for {
		percent = 0
		_, err := fmt.Scan(&n,&a)
		if err ==io.EOF {break}
		solution(n,a,1)
		fmt.Printf("%.5f", percent)
	}
}

// n 是剩余可以递归的次数，a是ai的值，p是当前递归层数的概率
func solution(n int, a int, p float64)  {
	if n < 1 {return}
	percent += 1.0/float64(a+1)*p
	for i:=1; i<=a; i++ {
		solution(n-1, i, p*1.0/float64(a+1))
	}
}