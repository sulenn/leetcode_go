package main

//https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/

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
	return mul(a, b)
}

func mul(a int, b int) int {
	result := 1
	for i := 0; i < a; i++ {
		result *= 3
		if result > 1000000007 {
			result %= 1000000007
		}
	}
	result *= b
	if result > 1000000007 {
		result %= 1000000007
	}
	return result
}
