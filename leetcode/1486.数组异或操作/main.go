package main

//https://leetcode-cn.com/problems/xor-operation-in-an-array/

func xorOperation(n int, start int) int {
	if n == 1 {
		return start
	}

	r := start

	for i := 1; i < n; i++ {
		start ^= r + 2*i
	}

	return start
}
