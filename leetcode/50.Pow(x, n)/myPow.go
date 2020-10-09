package main

import "fmt"

//https://leetcode-cn.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	flag := true
	if x < 0 {
		x = -x
		if n%2 != 0 {
			flag = false
		}
	}
	result := pow(x, n)
	if !flag {
		return -result
	}
	return result
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	var result float64
	result = pow(x, n/2)
	if n%2 == 1 {
		return result * result * x
	}
	return result * result
}

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
}
