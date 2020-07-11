package main

//https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/

//动态规划
func fib(n int) int { // 注意 n 取值为 1 - 100。所以考虑值溢出的问题，即累加值超过 math.maxint64 值
	if n < 2 {
		return n
	}
	fn_2 := 0 // f(n-2)
	fn_1 := 1 // f(n-1)
	for i := 2; i <= n; i++ {
		fn_1, fn_2 = fn_1+fn_2, fn_1
		if fn_1 > 1000000007 { // 注意每当 a + b 大于 1000000007 时都要取余。不然累加之后可能会溢出，大于 math.maxint64 值
			fn_1 %= 1000000007
		}
	}
	return fn_1
}
