package main

import "math"

//https://leetcode-cn.com/problems/sum-of-square-numbers/

//双指针
func judgeSquareSum(c int) bool {
	p1 := 0                          // 指针1
	p2 := int(math.Sqrt(float64(c))) // 指针2
	for p1 <= p2 {
		if p1*p1 > math.MaxInt64-p2*p2 {
			return false
		} // 溢出
		if p1*p1+p2*p2 == c {
			break
		}
		if p1*p1+p2*p2 > c {
			p2--
		} else {
			p1++
		}
	}
	return p1*p1+p2*p2 == c
}
