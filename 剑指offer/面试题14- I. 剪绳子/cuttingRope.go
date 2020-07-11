package main

import "math"

//https://leetcode-cn.com/problems/jian-sheng-zi-lcof/

//**参考**：https://leetcode-cn.com/problems/integer-break/solution/343-zheng-shu-chai-fen-tan-xin-by-jyd/
//贪心
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a := n / 3
	b := n % 3
	if b == 1 {
		b = 4
		a--
	}
	if b == 0 {
		b = 1
	}
	return int(math.Pow(float64(3), float64(a))) * b
}
