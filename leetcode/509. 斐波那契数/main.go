package main

//https://leetcode-cn.com/problems/fibonacci-number/

func fib(n int) int {
	if n < 2 {
		return 0
	}
	f0 := 0
	f1 := 1
	for i := 2; i < n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}
