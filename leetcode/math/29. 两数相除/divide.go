package main

import "math"

//https://leetcode-cn.com/problems/divide-two-integers/

//参考：https://leetcode-cn.com/problems/divide-two-integers/solution/xiao-xue-sheng-du-hui-de-lie-shu-shi-suan-chu-fa-b/
func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return math.MaxInt32
	} // int32 溢出
	flag := true
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}
	if dividend < 0 {
		flag = false
		dividend = -dividend
	}
	if divisor < 0 {
		flag = false
		divisor = -divisor
	}
	var count uint = 0
	for dividend >= divisor { // 求余数的二进制长度 + 1
		count++
		divisor = divisor << 1
	}
	result := 0
	for count > 0 { // 求余数
		count--
		divisor = divisor >> 1
		if dividend >= divisor {
			result += 1 << count
			dividend -= divisor
		}
	}
	if !flag {
		return -result
	}
	return result
}
