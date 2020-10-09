package main

//https://leetcode-cn.com/problems/sqrtx/

func mySqrt(x int) int {
	l := 1
	r := x
	for l <= r {
		mid := l + (r-l)/2
		if x/mid == mid {
			return mid
		}
		if x/mid > mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
