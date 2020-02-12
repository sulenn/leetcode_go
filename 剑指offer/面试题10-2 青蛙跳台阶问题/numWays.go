package main

//https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

//该题是 “斐波那契数列” 的变种，解法和 “斐波那契数列”解法一样
//注意 n 为 0 时输出为 1；n 为 1 输出为1； n 为 2 输出为 2；之后就是前两项累加
//f(n) = f(n-1) + f(n-2)
func numWays(n int) int {
	if n < 2 {return n}
	fn_2 := 0  // f(n-2)
	fn_1 := 1  // f(n-1)
	for i:=2; i<=n; i++ {
		fn_1, fn_2 = fn_1+fn_2, fn_1
		if fn_1 > 1000000007 { // 注意每当 a + b 大于 1000000007 时都要取余。不然累加之后可能会溢出，大于 math.maxint64 值
			fn_1 %= 1000000007
		}
	}
	return fn_1
}