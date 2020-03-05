package main

//https://leetcode-cn.com/problems/qiu-12n-lcof/

// 参考网友
func sumNums(n int) int {
	var f func(ret *int, n int) bool
	f = func(ret *int, n int) bool {
		*ret += n
		return n != 0 && f(ret, n-1)
	}
	var ret int
	f(&ret, n)
	return ret
}
