package main

//https://leetcode-cn.com/problems/qiu-12n-lcof/

//参考（逻辑运算符的短路效应）：https://leetcode-cn.com/problems/qiu-12n-lcof/solution/mian-shi-ti-64-qiu-1-2-nluo-ji-fu-duan-lu-qing-xi-/
func sumNums(n int) int {
	res := 0
	var recursive func(n int) int
	recursive = func(n int) int {
		_ = n > 1 && recursive(n-1) > 0
		res += n
		return res
	}
	recursive(n)
	return res
}
