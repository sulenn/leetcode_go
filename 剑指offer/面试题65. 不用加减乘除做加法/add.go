package main

//https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/

func add(a int, b int) int {
	sum := a ^ b  // 异或的值
	carry := (a & b) << 1   // 进位
	for carry != 0 {
		a = sum
		b = carry
		sum = a ^ b  // 异或的值
		carry = (a & b) << 1   // 进位
	}
	return sum | carry
}

